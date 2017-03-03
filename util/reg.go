package util

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// RegExport exports the windows registry key to the {regKey}.reg file.
func RegExport(regKey, root string) error {
	fileName := filepath.Join(root, filepath.Base(regKey)+".reg")
	regKey = getRegPath(regKey)

	cmd := exec.Command("reg", "export", regKey, fileName)
	return cmd.Run()
}

// If the windows architecture is x86, delete the `WOW6432Node\` from the registry key.
func getRegPath(regKey string) string {
	if runtime.GOARCH == "amd64" {
		return regKey
	}

	return strings.Replace(regKey, `WOW6432Node\`, "", 0)
}
