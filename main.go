package main

import (
	"flag"

	"log"

	"path/filepath"

	"github.com/kemokemo/wisloc/config"
	"github.com/kemokemo/wisloc/util"
)

func main() {
	infoPath := flag.String("i", "info.xml", "file path of the information xml to collect logs")
	outPath := flag.String("o", `.\`, "root path to save logs.")
	flag.Parse()

	rootpath := filepath.Clean(*outPath)

	// make directory to save files
	destPath, err := util.CreateUniqueDir(rootpath)
	if err != nil {
		log.Fatal("Failed to create the destination directory.", err)
	}

	// read xml file
	conf, err := config.LoadConfig(*infoPath)

	// save windows event logs
	if conf.IsNeedWindowsEventLogs {
		dst := filepath.Join(destPath, `eventlogs`)
		err = util.CheckAndMakeDir(dst)
		if err != nil {
			log.Fatal("Failed to create a directory.", err)
		}

		err = util.SaveEventLog(dst)
		if err != nil {
			log.Fatal("Failed to save the windows event logs.", err)
		}
	}

	// save application logs
	for _, item := range conf.LogPathInfoList {
		dst := filepath.Join(destPath, filepath.Base(item.Path))
		err = util.Copy(item.Path, dst)
		if err != nil {
			log.Fatal("Failed to copy.", err)
		}
	}

	// save application registry entries
	for _, reg := range conf.RegistryInfoList {
		err = util.RegExport(reg.Key, destPath)
		if err != nil {
			log.Fatal("Failed to export registry.", err)
		}
	}

	// archive the directory
	archiver := util.ZIP
	filename := archiver.DestFmt()(filepath.Base(destPath))
	err = archiver.Archive(destPath, filepath.Join(rootpath, filename))
	if err != nil {
		log.Println("Failed to archive.", err)
	}
}
