package pkg

import (
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//Server info
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Client information
type Client struct {
	Domains []string `yaml:"domains"`
	Token   string   `yaml:"token"`
}

//Config for client
type Config struct {
	Server   *Server `yaml:"server"`
	Token    string  `yaml:"token"`
	Duration string  `yaml:"duration"`
}

// ServerConfig config information for server
type ServerConfig struct {
	Port    int      `yaml:"port"`
	Clients []Client `yaml:"clients"`
}

// LoadConfig read config file
func LoadConfig(fileName string) *Config {
	p, err := filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(p)
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

// LoadServerConfig read server config from file
func LoadServerConfig(fileName string) *ServerConfig {

	cfg := &ServerConfig{}
	return cfg
}
