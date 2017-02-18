package util

import (
	"os/exec"
	"path/filepath"
)

const localhost = "127.0.0.1"

// SaveEventLog saves the windows event logs in both binary and text formats.
func SaveEventLog(dst string) error {
	logNames := []string{"Application", "System", "Security"}

	for _, name := range logNames {
		err := saveBinary(dst, name)
		if err != nil {
			return err
		}
	}

	for _, name := range logNames {
		err := saveText(dst, name)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveBinary(dst, name string) error {
	dst = filepath.Join(`.\`, dst)
	cmd := exec.Command("cscript", `scripts\WinEventLogBinary.vbs`, dst, name)
	return cmd.Run()
}

func saveText(dst, name string) error {
	filePath := filepath.Join(dst, name+".csv")
	cmd := exec.Command("cscript", `scripts\WinEventLogText.vbs`, localhost, filePath, name)
	return cmd.Run()
}
