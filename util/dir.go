package util

import (
	"os"
	"path/filepath"
	"time"
)

// CreateUniqueDir creates the unique name directory and return the directory path.
// The name is "{Hostname}_{Day-Time}" format.
func CreateUniqueDir(root string) (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return name, err
	}
	name += time.Now().Format("_20060102-030405")

	root = filepath.Clean(root)
	dest := filepath.Join(root, name)
	err = os.MkdirAll(dest, os.FileMode(777))
	if err != nil {
		return dest, err
	}

	return dest, nil
}

// CheckAndMakeDir checks if the destination directory exists.
// If there is no destination directory, it will create it.
func CheckAndMakeDir(dst string) error {
	dst = filepath.Clean(dst)
	_, err := os.Stat(dst)
	if err == nil {
		// the destination directory already exists
		return nil
	}

	err = os.Mkdir(dst, os.FileMode(777))
	if err != nil {
		return err
	}

	return nil
}
