package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	ApiHost string
	LBDir   string
	Msg     Message
}

var cfg config = config{
	ApiHost: "https://api.lstbknd.net",
	Msg:     Messages,
}

// Get config struct.
func Get() *config {

	return &cfg

}

// Set config from configPath. Type of file: yaml, yml
func Set(configPath string) error {

	var err error

	// Parsing config file
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		return err
	}

	return nil

}
