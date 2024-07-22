package credentials

import (
	"fmt"
	"sync"
)

type CredentialSource interface {
	FetchCredentials(registry string) (*Credential, error)
}

type Credential struct {
	username string
	password string
}

type Credentials struct {
	handlers map[string]CredentialSource
	l        sync.RWMutex
}

func NewCredentials() *Credentials {
	creds := &Credentials{
		handlers: make(map[string]CredentialSource),
	}
	return creds
}

func (creds *Credentials) RegisterSource(identifier string, handler CredentialSource) error {
	if identifier == "" {
		return fmt.Errorf("identifier must not be empty string")
	}
	if handler == nil {
		return fmt.Errorf("handler must not be nil")
	}
	creds.l.Lock()
	defer creds.l.Unlock()
	_, ok := creds.handlers[identifier]
	if ok {
		return fmt.Errorf("credential source %s already registered", identifier)
	}
	creds.handlers[identifier] = handler
	return nil
}

func (creds *Credentials) Source(identifier string) (CredentialSource, error) {
	creds.l.RLock()
	defer creds.l.RUnlock()
	h, ok := creds.handlers[identifier]
	if ok {
		return h, nil
	}
	return nil, fmt.Errorf("credential source %s not found", identifier)
}

func (creds *Credentials) FetchCredentials(identifier string, registry string) (*Credential, error) {
	src, err := creds.Source(identifier)
	if err != nil {
		return nil, err
	}
	return src.FetchCredentials(registry)
}
