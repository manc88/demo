package userservice

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CacheUsersListKey string `yaml:"cache_user_list"`
	UserCreationTopic string `yaml:"user_creation_topic"`
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
