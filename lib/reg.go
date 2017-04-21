package wisloc

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// RegExport exports the windows registry key to the {regKey}.reg file.
func RegExport(key, root string) error {
	path := filepath.Join(root, filepath.Base(key)+".txt")
	key = getRegPath(key)
	cmd := exec.Command("reg", "export", key, path)
	return cmd.Run()
}

// If the windows architecture is x86, delete the `WOW6432Node\` from the registry key.
func getRegPath(regKey string) string {
	if runtime.GOARCH == "amd64" {
		return regKey
	}
	return strings.Replace(regKey, `WOW6432Node\`, "", 0)
}
