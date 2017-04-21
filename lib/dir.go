package wisloc

import (
	"os"
	"path/filepath"
	"time"
)

// CreatePathInfo creates the unique directory path using root path.
// ex) "root\{Hostname}_{Day-Time}"
func CreatePathInfo(root string) (string, error) {
	h, err := os.Hostname()
	if err != nil {
		return "", err
	}
	dt := time.Now().Format("_20060102-030405")
	root = filepath.Clean(root)
	dest := filepath.Join(root, h+dt)

	return dest, nil
}
