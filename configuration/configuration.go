package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type WixConfiguration struct {
	Url       string `yaml:"url"`
	Auth      string `yaml:"auth"`
	AccountID string `yaml:"accountId"`
	SiteID    string `yaml:"siteId"`
}

type Configuration struct {
	WixConfig WixConfiguration `yaml:"wix"`
}

var Config Configuration

func Init() {
	log.Println("Loading config.yaml")
	yamlFile, err := ioutil.ReadFile("C:/Users/Stephen Balian/Desktop/2022-dev-projects/production-apps/escpos/config.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
