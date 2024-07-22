package credentials

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	argoexec "github.com/argoproj/pkg/exec"
)

type extScript struct {
	path string
	args []string
}

func (ext *extScript) FetchCredentials(registry string) (*Credential, error) {
	if !strings.HasPrefix(ext.path, "/") {
		return nil, fmt.Errorf("path to script must be absolute, but is '%s'", ext.path)
	}
	fi, err := os.Stat(ext.path)
	if err != nil {
		return nil, fmt.Errorf("could not stat %s: %v", ext.path, err)
	}
	if fi.Mode()&0111 == 0 {
		return nil, fmt.Errorf("script %s is not executable", ext.path)
	}
	cmd := exec.Command(ext.path, ext.args...)
	out, err := argoexec.RunCommandExt(cmd, argoexec.CmdOpts{Timeout: 10 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("error executing %s: %v", ext.path, err)
	}
	tokens := strings.SplitN(out, ":", 2)
	if len(tokens) != 2 {
		return nil, fmt.Errorf("invalid script output, must be single line with syntax <username>:<password>")
	}
	return &Credential{username: tokens[0], password: tokens[1]}, nil
}

func NewCredentialSourceExt(script string, args ...string) CredentialSource {
	return &extScript{path: script, args: args}
}
