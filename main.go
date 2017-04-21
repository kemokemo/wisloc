package main

import (
	"flag"
	"os"

	wisloc "github.com/kemokemo/wisloc/lib"

	"log"

	"path/filepath"
)

var (
	in   = flag.String("i", "info.xml", "file path of the information xml to collect logs")
	out  = flag.String("o", `.\`, "root path to save logs.")
	dest = ""
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	root, err := wisloc.CreatePathInfo(*out)
	if err != nil {
		log.Println(err)
		return 1
	}

	err = createPath(root)
	if err != nil {
		log.Println(err)
		return 1
	}

	c, err := wisloc.LoadConfig(*in)
	if err != nil {
		log.Printf("Failed to load config file:%s. %q", *in, err)
		return 1
	}

	err = collectWEL(root, c)
	if err != nil {
		log.Println("Failed to collect windows event logs.", err)
		return 1
	}

	err = collectFiles(root, c)
	if err != nil {
		log.Println("Failed to copy files and directories.", err)
		return 1
	}

	err = collectRegs(root, c)
	if err != nil {
		log.Println("Failed to export registry keys.", err)
		return 1
	}

	err = archive(root)
	if err != nil {
		log.Println("Failed to archive.", err)
		return 1
	}
	return 0
}

func createPath(path string) error {
	err := os.MkdirAll(path, os.FileMode(777))
	if err != nil {
		return err
	}
	return nil
}

func collectWEL(root string, c wisloc.Config) error {
	if c.NeedWindowsEventLog == false {
		return nil
	}
	dst := filepath.Join(root, `eventlogs`)
	err := createPath(dst)
	if err != nil {
		return err
	}
	return wisloc.CollectEventLogs(dst)
}

func collectFiles(root string, c wisloc.Config) error {
	dst := filepath.Join(root, `logfiles`)
	err := createPath(dst)
	if err != nil {
		return err
	}
	for _, item := range c.PathInfoList {
		dt := filepath.Join(dst, filepath.Base(item.Path))
		err = wisloc.Copy(item.Path, dt)
		if err != nil {
			return err
		}
	}
	return nil
}

func collectRegs(root string, c wisloc.Config) error {
	dst := filepath.Join(root, `registries`)
	err := createPath(dst)
	if err != nil {
		return err
	}
	for _, reg := range c.RegInfoList {
		err := wisloc.RegExport(reg.Key, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func archive(root string) error {
	archiver := wisloc.ZIP
	filename := archiver.DestFmt()(filepath.Base(root))
	return archiver.Archive(root, filepath.Join(".", filename))
}
