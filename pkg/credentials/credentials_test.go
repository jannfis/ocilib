package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_NewCredentials(t *testing.T) {
	c := NewCredentials()
	assert.NotNil(t, c)
}

func Test_RegisterSources(t *testing.T) {
	t.Run("Register a source", func(t *testing.T) {
		c := NewCredentials()
		credSrc := NewMockCredentialSource(t)
		err := c.RegisterSource("foo", credSrc)
		assert.NoError(t, err)
		err = c.RegisterSource("foo", credSrc)
		assert.Error(t, err)
	})
}

func Test_FetchCredentials(t *testing.T) {
	t.Run("Fetch a credential", func(t *testing.T) {
		c := NewCredentials()
		credSrc := NewMockCredentialSource(t)
		credSrc.On("FetchCredentials", mock.Anything).Return(&Credential{username: "user", password: "pass"}, nil)
		err := c.RegisterSource("mock", credSrc)
		assert.NoError(t, err)
		creds, err := c.FetchCredentials("mock", "gcr.io")
		require.NoError(t, err)
		assert.Equal(t, "user", creds.username)
		assert.Equal(t, "pass", creds.password)
	})
}
