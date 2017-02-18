// Create a unique directory using the hostname and date information.
package util

import (
	"os"
	"path/filepath"
	"time"
)

// CreateUniqueDir creates the unique name directory.
// The name is "{Hostname}_{Day-Time}" format.
func CreateUniqueDir() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return name, err
	}

	name += time.Now().Format("_20060102-030405")
	err = os.Mkdir(name, os.FileMode(777))
	if err != nil {
		return name, err
	}

	return name, nil
}

// CreateDstPath creates the destination path from root path and src path.
func CreateDstPath(rootPath, src string) (dst string) {
	src = filepath.Clean(src)
	dst = filepath.Join(rootPath, filepath.Base(src))
	return dst
}

// If the destination directory does not exist, CheckAndCreateDir creates
// the directory.
func CheckAndCreateDir(dst string) error {
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
