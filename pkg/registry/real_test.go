package registry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Digest(t *testing.T) {
	t.Run("Get a digest", func(t *testing.T) {
		ep, err := GetRegistryEndpoint("quay.io")
		require.NoError(t, err)
		require.NotNil(t, ep)
	})
}
