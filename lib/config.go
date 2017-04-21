package wisloc

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Config is the configuration information to collect logs.
type Config struct {
	XMLName             xml.Name   `xml:"CollectSettings"`
	SoftName            string     `xml:"SoftName"`
	NeedWindowsEventLog bool       `xml:"NeedWEL"`
	RegInfoList         []RegInfo  `xml:"RegInfoList>RegInfo"`
	PathInfoList        []PathInfo `xml:"PathInfoList>PathInfo"`
}

// RegInfo is the registry information to collect.
type RegInfo struct {
	XMLName xml.Name `xml:"RegInfo"`
	Key     string   `xml:"Key"`
}

// PathInfo is the directory and file paths information to collect.
type PathInfo struct {
	XMLName xml.Name `xml:"PathInfo"`
	Path    string   `xml:"Path"`
}

// LoadConfig loads the Config struct from xml file specified an arg.
func LoadConfig(path string) (Config, error) {
	config := Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = xml.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, err
	}
	for i, info := range config.PathInfoList {
		config.PathInfoList[i].Path = convertEnv(info.Path)
	}
	return config, nil
}

// convertEnv converts the Windows environment variables contained path.
func convertEnv(path string) string {
	var converted string
	for _, p := range strings.Split(path, "%") {
		value := os.Getenv(p)
		if value == "" {
			value = p
		}
		converted = filepath.Join(converted, value)
	}
	return converted
}
