package grpctransport

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(fileName string) (*Config, error) {
	cfg := Config{}

	yamlBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
