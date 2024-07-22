package credentials

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const magic = "# ocilib-1.0 file store"

type file struct {
	path string
}

func (f *file) FetchCredentials(registry string) (*Credential, error) {
	fd, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)

	// Read the first line and check for magic string
	if !scanner.Scan() || scanner.Text() != magic {
		return nil, fmt.Errorf("%s: magic not found", f.path)
	}

	lineNo := 1

	// We scan through the whole file before returning the credentials, to
	// make sure we have a completely valid credentials file.
	for scanner.Scan() {
		lineNo += 1
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		tokens := strings.SplitN(scanner.Text(), ":", 3)
		if len(tokens) != 3 || tokens[0] == "" || tokens[1] == "" || tokens[2] == "" {
			return nil, fmt.Errorf("%s:%d: malformed data", f.path, lineNo)
		}
		if tokens[0] != registry {
			continue
		}
		return &Credential{username: tokens[1], password: tokens[2]}, nil
	}

	return nil, fmt.Errorf("no credentials found for %s", registry)
}

// NewCredentialSourceFile creates a new file backed credential source from the
// file at given path.
//
// File credentials supports simple username/password credentials per registry,
// one entry per line.
//
// The first line in the file MUST read the following:
//
// # ocilib-1.0 file store
//
// Lines in the file must have the following format:
//
//	<registry>:<username>:<password>
//
// Empty lines and lines starting with a hash character ('#') will be ignored.
func NewCredentialSourceFile(path string) CredentialSource {
	return &file{path: path}
}
