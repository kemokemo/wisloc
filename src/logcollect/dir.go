// Create a unique directory using the hostname and date information.
package logcollect

import (
	"os"
	"log"
	"time"
)

func MakeUniqueDir() {
	name, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get the hostname")
	}

	t := time.Now()
	name += t.Format("_20060102-030405")
	os.Mkdir(name, os.FileMode(777))
}