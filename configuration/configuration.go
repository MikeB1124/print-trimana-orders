package configuration

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/MikeB1124/print-trimana-orders/logger"
	"gopkg.in/yaml.v3"
)

type WixConfiguration struct {
	Url       string `yaml:"url"`
	Auth      string `yaml:"auth"`
	AccountID string `yaml:"accountId"`
	SiteID    string `yaml:"siteId"`
}

type Printer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Configuration struct {
	WixConfig WixConfiguration `yaml:"wix"`
	Printers  []Printer        `yaml:"printers"`
}

var Config Configuration

func Init() {
	var configPath string
	if os.Getenv("CONTAINER") == "true" {
		configPath = "/home/config.yaml"
	} else {
		configPath = "config.yaml"
	}
	logger.InfoLogger.Printf("Config path set to %s\n", configPath)
	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
