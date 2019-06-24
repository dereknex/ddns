package config

import "sync"

type DomainConfig struct {
	Provider string `yaml:"provider"`
	Entrys []string `yaml:"entrys"`
}

type clientConfig struct {
	Solvers []string            `yaml:"solvers"`
	Domains map[string]DomainConfig `yaml:"domains"`
}


var ClientConfig *clientConfig;
var newConfig sync.Once

func init(){
	newConfig.Do(func() {
		ClientConfig = &clientConfig{}
	})
}