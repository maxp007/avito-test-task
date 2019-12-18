package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var conf *Config

func GetInstance() *Config {
	return conf
}

type Config struct {
	Data *ConfStruct
}

func init() {
	log.SetReportCaller(true)
	conf = &Config{&ConfStruct{}}
}

func (c *Config) LoadConfig() (err error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return
	}

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(c.Data)
	if err != nil {
		return
	}

	err = f.Close()
	if err != nil {
		return
	}
	return
}
