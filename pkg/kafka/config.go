package kafka

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Brokers     []string `yaml:"brokers"`
	MaxAttempts int      `yaml:"max_attempts"`
	Async       bool     `yaml:"async"`
}

func NewConfig(file string) (*Config, error) {
	cfg := Config{}

	yamlBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
