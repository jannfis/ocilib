package registry

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/argoproj/pkg/json"

	"github.com/jannfis/ocilib/pkg/options"
	"github.com/jannfis/ocilib/pkg/tag"

	"github.com/jannfis/ocilib/pkg/log"
	"github.com/jannfis/ocilib/pkg/metrics"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest/manifestlist"
	"github.com/distribution/distribution/v3/manifest/ocischema"
	"github.com/distribution/distribution/v3/manifest/schema1" //nolint:staticcheck
	"github.com/distribution/distribution/v3/manifest/schema2"
	"github.com/distribution/distribution/v3/reference"
	"github.com/distribution/distribution/v3/registry/client"
	"github.com/distribution/distribution/v3/registry/client/auth"
	"github.com/distribution/distribution/v3/registry/client/auth/challenge"
	"github.com/distribution/distribution/v3/registry/client/transport"

	"github.com/opencontainers/go-digest"
	ociv1 "github.com/opencontainers/image-spec/specs-go/v1"

	"go.uber.org/ratelimit"

	"net/http"
	"net/url"
	"strings"
	"time"
)

// TODO: Check image's architecture and OS

// knownMediaTypes is the list of media types we can process
var knownMediaTypes = []string{
	ocischema.SchemaVersion.MediaType,
	schema1.MediaTypeSignedManifest, //nolint:staticcheck
	schema2.SchemaVersion.MediaType,
	manifestlist.SchemaVersion.MediaType,
	ociv1.MediaTypeImageIndex,
}

// RegistryClient defines the methods we need for querying container registries
type RegistryClient interface {
	NewRepository(nameInRepository string) error
	Tags() ([]string, error)
	ManifestForTag(tagStr string) (distribution.Manifest, error)
	ManifestForDigest(dgst digest.Digest) (distribution.Manifest, error)
	TagMetadata(manifest distribution.Manifest, opts *options.ManifestOptions) (*tag.TagInfo, error)
	TagInfoFromReferences(opts *options.ManifestOptions, ti *tag.TagInfo, references []distribution.Descriptor) (*tag.TagInfo, error)
	// Repository() distribution.Repository
}

type NewRegistryClient func(*RegistryEndpoint, string, string) (RegistryClient, error)

// Helper type for registry clients
type registryClient struct {
	regClient distribution.Repository
	endpoint  *RegistryEndpoint
	creds     *credentials
}

// credentials is an implementation of distribution/V3/session struct
// to mangage registry credentials and token
type credentials struct {
	username      string
	password      string
	refreshTokens map[string]string
	l             sync.RWMutex
}

func (c *credentials) Basic(url *url.URL) (string, string) {
	return c.username, c.password
}

func (c *credentials) RefreshToken(url *url.URL, service string) string {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.refreshTokens[service]
}

func (c *credentials) SetRefreshToken(realm *url.URL, service, token string) {
	c.l.Lock()
	defer c.l.RUnlock()
	if c.refreshTokens != nil {
		c.refreshTokens[service] = token
	}
}

// rateLimitTransport encapsulates our custom HTTP round tripper with rate
// limiter from the endpoint.
type rateLimitTransport struct {
	limiter   ratelimit.Limiter
	transport http.RoundTripper
	endpoint  *RegistryEndpoint
}

// RoundTrip is a custom RoundTrip method with rate-limiter
func (rlt *rateLimitTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	rlt.limiter.Take()
	log.Tracef("Performing HTTP %s %s", r.Method, r.URL)
	resp, err := rlt.transport.RoundTrip(r)
	metrics.Endpoint().IncreaseRequest(rlt.endpoint.RegistryAPI, err != nil)
	return resp, err
}

// NewRepository is a wrapper for creating a registry client that is possibly
// rate-limited by using a custom HTTP round tripper method.
func (clt *registryClient) NewRepository(nameInRepository string) error {
	urlToCall := strings.TrimSuffix(clt.endpoint.RegistryAPI, "/")
	challengeManager1 := challenge.NewSimpleManager()
	_, err := ping(challengeManager1, clt.endpoint, "")
	if err != nil {
		return err
	}

	authTransport := transport.NewTransport(
		clt.endpoint.GetTransport(), auth.NewAuthorizer(
			challengeManager1,
			auth.NewTokenHandler(clt.endpoint.GetTransport(), clt.creds, nameInRepository, "pull"),
			auth.NewBasicHandler(clt.creds)))

	rlt := &rateLimitTransport{
		limiter:   clt.endpoint.Limiter,
		transport: authTransport,
		endpoint:  clt.endpoint,
	}

	named, err := reference.WithName(nameInRepository)
	if err != nil {
		return err
	}
	clt.regClient, err = client.NewRepository(named, urlToCall, rlt)
	if err != nil {
		return err
	}
	return nil
}

// NewClient returns a new RegistryClient for the given endpoint information
func NewClient(endpoint *RegistryEndpoint, username, password string) (RegistryClient, error) {
	if username == "" && endpoint.Username != "" {
		username = endpoint.Username
	}
	if password == "" && endpoint.Password != "" {
		password = endpoint.Password
	}
	creds := &credentials{
		username: username,
		password: password,
	}
	return &registryClient{
		creds:    creds,
		endpoint: endpoint,
	}, nil
}

// So that we can set a potential mocked repository client
func (clt *registryClient) setDistributionRepository(repo distribution.Repository) {
	clt.regClient = repo
}

// Tags returns a list of tags for given name in repository
func (clt *registryClient) Tags() ([]string, error) {
	tagService := clt.regClient.Tags(context.Background())
	tTags, err := tagService.All(context.Background())
	if err != nil {
		return nil, err
	}
	return tTags, nil
}

// Manifest  returns a Manifest for a given tag in repository
func (clt *registryClient) ManifestForTag(tagStr string) (distribution.Manifest, error) {
	manService, err := clt.regClient.Manifests(context.Background())
	if err != nil {
		return nil, err
	}
	manifest, err := manService.Get(
		context.Background(),
		digest.FromString(tagStr),
		distribution.WithTag(tagStr), distribution.WithManifestMediaTypes(knownMediaTypes))
	if err != nil {
		return nil, err
	}
	return manifest, nil
}

// ManifestForDigest  returns a Manifest for a given digest in repository
func (clt *registryClient) ManifestForDigest(dgst digest.Digest) (distribution.Manifest, error) {
	manService, err := clt.regClient.Manifests(context.Background())
	if err != nil {
		return nil, err
	}
	manifest, err := manService.Get(
		context.Background(),
		dgst,
		distribution.WithManifestMediaTypes(knownMediaTypes))
	if err != nil {
		return nil, err
	}
	return manifest, nil
}

// TagMetadata retrieves metadata for a given manifest of given repository
func (client *registryClient) TagMetadata(manifest distribution.Manifest, opts *options.ManifestOptions) (*tag.TagInfo, error) {
	ti := &tag.TagInfo{}
	logCtx := opts.Logger()
	var info struct {
		Arch    string `json:"architecture"`
		Created string `json:"created"`
		OS      string `json:"os"`
		Variant string `json:"variant"`
	}

	// We support the following types of manifests as returned by the registry:
	//
	// V1 (legacy, might go away), V2 and OCI
	//
	// Also ManifestLists (e.g. on multi-arch images) are supported.
	//
	switch deserialized := manifest.(type) {

	case *schema1.SignedManifest: //nolint:staticcheck
		var man schema1.Manifest = deserialized.Manifest //nolint:staticcheck
		if len(man.History) == 0 {
			return nil, fmt.Errorf("no history information found in schema V1")
		}

		_, mBytes, err := manifest.Payload()
		if err != nil {
			return nil, err
		}
		ti.Digest = sha256.Sum256(mBytes)

		logCtx.Tracef("v1 SHA digest is %s", ti.EncodedDigest())
		if err := json.Unmarshal([]byte(man.History[0].V1Compatibility), &info); err != nil {
			return nil, err
		}
		if !opts.WantsPlatform(info.OS, info.Arch, "") {
			logCtx.Debugf("ignoring v1 manifest %v. Manifest platform: %s, requested: %s",
				ti.EncodedDigest(), options.PlatformKey(info.OS, info.Arch, info.Variant), strings.Join(opts.Platforms(), ","))
			return nil, nil
		}
		if createdAt, err := time.Parse(time.RFC3339Nano, info.Created); err != nil {
			return nil, err
		} else {
			ti.CreatedAt = createdAt
		}
		return ti, nil

	case *manifestlist.DeserializedManifestList:
		var list manifestlist.DeserializedManifestList = *deserialized

		// List must contain at least one image manifest
		if len(list.Manifests) == 0 {
			return nil, fmt.Errorf("empty manifestlist not supported")
		}

		// We use the SHA from the manifest list to let the container engine
		// decide which image to pull, in case of multi-arch clusters.
		_, mBytes, err := list.Payload()
		if err != nil {
			return nil, fmt.Errorf("could not retrieve manifestlist payload: %v", err)
		}
		ti.Digest = sha256.Sum256(mBytes)

		logCtx.Tracef("SHA256 of manifest parent is %v", ti.EncodedDigest())

		return client.TagInfoFromReferences(opts, ti, list.References())

	case *ocischema.DeserializedImageIndex:
		var index ocischema.DeserializedImageIndex = *deserialized

		// Index must contain at least one image manifest
		if len(index.Manifests) == 0 {
			return nil, fmt.Errorf("empty index not supported")
		}

		// We use the SHA from the manifest index to let the container engine
		// decide which image to pull, in case of multi-arch clusters.
		_, mBytes, err := index.Payload()
		if err != nil {
			return nil, fmt.Errorf("could not retrieve index payload: %v", err)
		}
		ti.Digest = sha256.Sum256(mBytes)

		logCtx.Tracef("SHA256 of manifest parent is %v", ti.EncodedDigest())

		return client.TagInfoFromReferences(opts, ti, index.References())

	case *schema2.DeserializedManifest:
		var man schema2.Manifest = deserialized.Manifest

		logCtx.Tracef("Manifest digest is %v", man.Config.Digest.Encoded())

		_, mBytes, err := manifest.Payload()
		if err != nil {
			return nil, err
		}
		ti.Digest = sha256.Sum256(mBytes)
		logCtx.Tracef("v2 SHA digest is %s", ti.EncodedDigest())

		// The data we require from a V2 manifest is in a blob that we need to
		// fetch from the registry.
		blobReader, err := client.regClient.Blobs(context.Background()).Get(context.Background(), man.Config.Digest)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(blobReader, &info); err != nil {
			return nil, err
		}

		if !opts.WantsPlatform(info.OS, info.Arch, info.Variant) {
			logCtx.Debugf("ignoring v2 manifest %v. Manifest platform: %s/%s, requested: %s",
				ti.EncodedDigest(), info.OS, info.Arch, strings.Join(opts.Platforms(), ","))
			return nil, nil
		}

		if ti.CreatedAt, err = time.Parse(time.RFC3339Nano, info.Created); err != nil {
			return nil, err
		}

		return ti, nil
	case *ocischema.DeserializedManifest:
		var man ocischema.Manifest = deserialized.Manifest

		_, mBytes, err := manifest.Payload()
		if err != nil {
			return nil, err
		}
		ti.Digest = sha256.Sum256(mBytes)
		logCtx.Tracef("OCI SHA digest is %s", ti.EncodedDigest())

		// The data we require from a V2 manifest is in a blob that we need to
		// fetch from the registry.
		blobReader, err := client.regClient.Blobs(context.Background()).Get(context.Background(), man.Config.Digest)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(blobReader, &info); err != nil {
			return nil, err
		}

		if !opts.WantsPlatform(info.OS, info.Arch, info.Variant) {
			logCtx.Debugf("ignoring OCI manifest %v. Manifest platform: %s/%s, requested: %s",
				ti.EncodedDigest(), info.OS, info.Arch, strings.Join(opts.Platforms(), ","))
			return nil, nil
		}

		if ti.CreatedAt, err = time.Parse(time.RFC3339Nano, info.Created); err != nil {
			return nil, err
		}

		return ti, nil
	default:
		return nil, fmt.Errorf("invalid manifest type %T", manifest)
	}
}

// TagInfoFromReferences is a helper method to retrieve metadata for a given
// list of references. It will return the most recent pushed manifest from the
// list of references.
func (client *registryClient) TagInfoFromReferences(opts *options.ManifestOptions, ti *tag.TagInfo, references []distribution.Descriptor) (*tag.TagInfo, error) {
	var ml []distribution.Descriptor
	platforms := []string{}
	logCtx := opts.Logger()

	for _, ref := range references {
		platforms = append(platforms, ref.Platform.OS+"/"+ref.Platform.Architecture)
		logCtx.Tracef("Found %s", options.PlatformKey(ref.Platform.OS, ref.Platform.Architecture, ref.Platform.Variant))
		if !opts.WantsPlatform(ref.Platform.OS, ref.Platform.Architecture, ref.Platform.Variant) {
			logCtx.Tracef("Ignoring referenced manifest %v because platform %s does not match any of: %s",
				ref.Digest,
				options.PlatformKey(ref.Platform.OS, ref.Platform.Architecture, ref.Platform.Variant),
				strings.Join(opts.Platforms(), ","))
			continue
		}
		ml = append(ml, ref)
	}

	// We need at least one reference that matches requested plaforms
	if len(ml) == 0 {
		logCtx.Debugf("Manifest list did not contain any usable reference. Platforms requested: (%s), platforms included: (%s)",
			strings.Join(opts.Platforms(), ","), strings.Join(platforms, ","))
		return nil, nil
	}

	// For some strategies, we do not need to fetch metadata for further
	// processing.
	if !opts.WantsMetadata() {
		return ti, nil
	}

	// Loop through all referenced manifests to get their metadata. We only
	// consider manifests for platforms we are interested in.
	for _, ref := range ml {
		logCtx.Tracef("Inspecting metadata of reference: %v", ref.Digest)

		man, err := client.ManifestForDigest(ref.Digest)
		if err != nil {
			return nil, fmt.Errorf("could not fetch manifest %v: %v", ref.Digest, err)
		}

		cti, err := client.TagMetadata(man, opts)
		if err != nil {
			return nil, fmt.Errorf("could not fetch metadata for manifest %v: %v", ref.Digest, err)
		}

		// We save the timestamp of the most recent pushed manifest for any
		// given reference, if the metadata for the tag was correctly
		// retrieved. This is important for the latest update strategy to
		// be able to handle multi-arch images. The latest strategy will
		// consider the most recent reference from an image index.
		if cti != nil {
			if cti.CreatedAt.After(ti.CreatedAt) {
				ti.CreatedAt = cti.CreatedAt
			}
		} else {
			logCtx.Warnf("returned metadata for manifest %v is nil, this should not happen.", ref.Digest)
			continue
		}
	}

	return ti, nil
}

// Implementation of ping method to intialize the challenge list
// Without this, tokenHandler and AuthorizationHandler won't work
func ping(manager challenge.Manager, endpoint *RegistryEndpoint, versionHeader string) ([]auth.APIVersion, error) {
	httpc := &http.Client{Transport: endpoint.GetTransport()}
	url := endpoint.RegistryAPI + "/v2/"
	resp, err := httpc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Let's consider only HTTP 200 and 401 valid responses for the inital request
	if resp.StatusCode != 200 && resp.StatusCode != 401 {
		return nil, fmt.Errorf("endpoint %s does not seem to be a valid v2 Docker Registry API (received HTTP code %d for GET %s)", endpoint.RegistryAPI, resp.StatusCode, url)
	}

	if err := manager.AddResponse(resp); err != nil {
		return nil, err
	}

	return auth.APIVersions(resp, versionHeader), err
}
