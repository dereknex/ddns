package config

type ClientConfig struct {
	Solvers []string            `yaml:"solvers"`
	Domains map[string][]string `yaml:"domains"`
}
