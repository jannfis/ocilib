package credentials

import (
	"fmt"
	"os"
	"strings"
)

type environment struct {
	envName string
}

func (e *environment) FetchCredentials(registry string) (*Credential, error) {
	env, ok := os.LookupEnv(e.envName)
	if !ok {
		return nil, fmt.Errorf("environment %s is not set", e.envName)
	}
	ss := strings.SplitN(env, ":", 2)
	if len(ss) != 2 || ss[0] == "" || ss[1] == "" {
		return nil, fmt.Errorf("malformed credentials in environment %s", e.envName)
	}
	return &Credential{username: ss[0], password: ss[1]}, nil
}

func NewCredentialSourceEnvironment(envName string) CredentialSource {
	return &environment{envName: envName}
}
