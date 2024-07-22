package fixture

// import (
// 	"context"

// 	"github.com/distribution/distribution/v3"
// 	"github.com/distribution/distribution/v3/registry/storage"
// 	"github.com/distribution/distribution/v3/registry/storage/driver/inmemory"
// 	"github.com/distribution/distribution/v3/registry"
// )

// func NewRegistry() distribution.Namespace {
// 	// b := MustReadFile("config.yaml")
// 	stg := inmemory.New()
// 	r, err := storage.NewRegistry(context.TODO(), stg)
// 	// cfg := &configuration.Configuration{Version: "0.1"}
// 	// p := configuration.NewParser("", []configuration.VersionedParseInfo{configuration.VersionedParseInfo{Version: "0.1"}})
// 	// err := p.Parse([]byte(b), cfg)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	rr, err := registry.NewRegistry(context.TODO(), cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return r
// }
