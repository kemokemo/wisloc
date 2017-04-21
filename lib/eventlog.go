package wisloc

import (
	"os/exec"
	"path/filepath"
)

const (
	localhost = "127.0.0.1"
	app       = "Application"
	sys       = "System"
)

// CollectEventLogs saves the windows event logs in both binary and text formats.
func CollectEventLogs(dst string) error {
	names := []string{app, sys}
	for _, name := range names {
		err := saveBinary(dst, name)
		if err != nil {
			return err
		}
		err = saveText(dst, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveBinary(dst, name string) error {
	dst = filepath.Join(`.\`, dst)
	cmd := exec.Command("cscript", `scripts\welb.vbs`, dst, name)
	return cmd.Run()
}

func saveText(dst, name string) error {
	filePath := filepath.Join(dst, name+".csv")
	cmd := exec.Command("cscript", `scripts\welt.vbs`, localhost, filePath, name)
	return cmd.Run()
}
