package pgx

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	Database           string `yaml:"database"`
	Sslmode            string `yaml:"sslmode"`
	MaxConnIdleTimeSec int    `yaml:"maxConnIdleTimeSec,omitempty"`
	MaxConnLifetimeSec int    `yaml:"maxConnLifetimeSec,omitempty"`
	MinConns           int32  `yaml:"minConns,omitempty"`
	MaxConns           int32  `yaml:"maxConns,omitempty"`
}

func NewConfigFromFile(file string) (*Config, error) {
	cfg := Config{
		MaxConnIdleTimeSec: 60,
		MaxConnLifetimeSec: 3600,
		MinConns:           2,
		MaxConns:           4,
	}

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
