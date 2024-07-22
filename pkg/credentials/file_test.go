package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CredentialSourceFile(t *testing.T) {
	t.Run("Non-existing path", func(t *testing.T) {
		cs := NewCredentialSourceFile("what")
		c, err := cs.FetchCredentials("ghcr.io")
		assert.ErrorIs(t, err, os.ErrNotExist)
		assert.Nil(t, c)
	})
	t.Run("Valid credential file", func(t *testing.T) {
		cs := NewCredentialSourceFile("testdata/file/valid.txt")

		c, err := cs.FetchCredentials("ghcr.io")
		require.NoError(t, err)
		assert.Equal(t, "foo", c.username)
		assert.Equal(t, "bar", c.password)

		c, err = cs.FetchCredentials("docker.io")
		require.NoError(t, err)
		assert.Equal(t, "foo", c.username)
		assert.Equal(t, "baz", c.password)

		c, err = cs.FetchCredentials("unconfigured")
		require.ErrorContains(t, err, "no credentials found")
		assert.Nil(t, c)
	})
	t.Run("Missing magic in credential file", func(t *testing.T) {
		cs := NewCredentialSourceFile("testdata/file/nomagic.txt")

		c, err := cs.FetchCredentials("ghcr.io")
		assert.ErrorContains(t, err, "magic not found")
		assert.Nil(t, c)
		c, err = cs.FetchCredentials("docker.io")
		assert.ErrorContains(t, err, "magic not found")
		assert.Nil(t, c)
	})

	t.Run("Malformed data in credential file", func(t *testing.T) {
		cs := NewCredentialSourceFile("testdata/file/malformed.txt")

		c, err := cs.FetchCredentials("ghcr.io")
		assert.ErrorContains(t, err, "malformed data")
		assert.Nil(t, c)
	})

}
