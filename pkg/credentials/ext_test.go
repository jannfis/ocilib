package credentials

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var myPath string

func Test_ExtCredentialSource(t *testing.T) {
	scriptPath := path.Join(myPath, "./testdata/ext/good.sh")
	cs := NewCredentialSourceExt(scriptPath)
	creds, err := cs.FetchCredentials("ghcr.io")
	require.NoError(t, err)
	assert.Equal(t, "foo", creds.username)
	assert.Equal(t, "bar", creds.password)
}

func init() {
	var err error
	myPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
