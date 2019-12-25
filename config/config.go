package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
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

	err := GetInstance().loadConfig()
	if err != nil {
		log.Fatal("config reader,", err)
	}

}

func (c *Config) loadConfig() (err error) {
	filename, err := filepath.Abs("./config/config.yaml")

	if err != nil {

		return
	}

	f, err := os.Open(filename)
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
