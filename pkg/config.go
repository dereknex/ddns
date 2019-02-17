package pkg

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//Server info
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

//Config for client
type Config struct {
	Server   *Server `yaml:"server"`
	Token    string  `yaml:"token"`
	Duration string  `yaml:"duration"`
}

// LoadConfig read config file
func LoadConfig(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	cfg := &Config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Fatalf("Read config error: %s", err.Error())
		return nil
	}
	return cfg
}
