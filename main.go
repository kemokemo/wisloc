package main

import (
	"flag"

	"log"

	"github.com/kemokemo/logcollector/config"
	"github.com/kemokemo/logcollector/util"
)

func main() {
	xmlpath := flag.String("xml", "info.xml", "information xml file path to collect logs")
	flag.Parse()

	// make directory to save files
	root, err := util.CreateUniqueDir()
	if err != nil {
		log.Fatal("Failed to create the destination directory.", err)
	}

	// read xml file
	conf, err := config.LoadConfig(*xmlpath)

	// save windows event logs
	if conf.IsNeedWindowsEventLogs {
		dst := util.CreateDstPath(root, `eventlogs`)
		err = util.CheckAndCreateDir(dst)
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
		dst := util.CreateDstPath(root, item.Path)
		err = util.Copy(item.Path, dst)
		if err != nil {
			log.Fatal("Failed to copy.", err)
		}
	}

	// save application registry entries
	for _, reg := range conf.RegistryInfoList {
		err = util.RegExport(reg.Key, root)
		if err != nil {
			log.Fatal("Failed to export registry.", err)
		}
	}

	// compress the directory
}
