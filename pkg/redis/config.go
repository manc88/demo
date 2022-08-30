package redis

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	TTL      int64  `yaml:"ttl"`
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
