package config

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

// Config is the configuration information to collect logs.
type Config struct {
	XMLName                xml.Name       `xml:"CollectingSettings"`
	SoftwareName           string         `xml:"SoftwareName"`
	IsNeedWindowsEventLogs bool           `xml:"IsNeedWindowsEventLogs"`
	RegistryInfoList       []RegistryInfo `xml:"RegistryInfoList>RegistryInfo"`
	LogPathInfoList        []LogPathInfo  `xml:"LogPathInfoList>LogPathInfo"`
}

// RegistryInfo is the registry information to collect.
type RegistryInfo struct {
	XMLName xml.Name `xml:"RegistryInfo"`
	Key     string   `xml:"Key"`
}

// LogPathInfo is the paths information to collect.
type LogPathInfo struct {
	XMLName xml.Name `xml:"LogPathInfo"`
	Path    string   `xml:"Path"`
}

// LoadConfig loads the Config struct from xml file.
func LoadConfig(filePath string) (Config, error) {
	config := Config{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	if err = xml.Unmarshal([]byte(data), &config); err != nil {
		return config, err
	}

	// 読み込んだLogPath情報に含まれる環境変数を変換する
	for i := range config.LogPathInfoList {
		var logPath string
		for _, p := range strings.Split(config.LogPathInfoList[i].Path, "%") {
			logPath += convertEnvInfo(p)
		}
		config.LogPathInfoList[i].Path = logPath
	}

	return config, nil
}

// If envInfo exist in the Windows environment variables, converted string will return.
// Otherwise, original string will return.
func convertEnvInfo(envInfo string) string {
	value := os.Getenv(envInfo)
	if value == "" {
		return envInfo
	}
	return value
}
