package wisloc

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Archiver is an interface to compress files and directories to archive.
type Archiver interface {
	DestFmt() func(string) string
	Archive(src, dest string) error
}

type zipper struct {
}

// ZIP is an Archiver to compress and decompress as zip file format
var ZIP Archiver = (*zipper)(nil)

// Archive function archives source to ZIP format archive.
func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // Skip
		}
		if err != nil {
			return err
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}

func (z *zipper) DestFmt() func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s.zip", name)
	}
}
