package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CredsFromEnvironment(t *testing.T) {
	t.Run("Get credentials from environment", func(t *testing.T) {
		tts := []struct {
			creds    string
			username string
			password string
		}{
			{"foo:bar", "foo", "bar"},
			{"foo:bar:baz", "foo", "bar:baz"},
			{"foo::::::", "foo", ":::::"},
		}
		for _, tt := range tts {
			t.Setenv("CREDENTIALS", tt.creds)
			ec := NewCredentialSourceEnvironment("CREDENTIALS")
			c, err := ec.FetchCredentials("gcr.io")
			assert.NoError(t, err)
			assert.Equal(t, tt.username, c.username)
			assert.Equal(t, tt.password, c.password)
		}
	})
	t.Run("Environment not set", func(t *testing.T) {
		ec := NewCredentialSourceEnvironment("CREDENTIALS")
		c, err := ec.FetchCredentials("gcr.io")
		assert.Error(t, err)
		assert.Nil(t, c)
	})
	t.Run("Malformed credentials", func(t *testing.T) {
		for _, creds := range []string{"foo", ":", "::", "foo:", ":bar"} {
			t.Setenv("CREDENTIALS", creds)
			ec := NewCredentialSourceEnvironment("CREDENTIALS")
			c, err := ec.FetchCredentials("gcr.io")
			assert.Error(t, err)
			assert.Nil(t, c)
		}
	})

}
