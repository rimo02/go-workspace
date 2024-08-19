package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type BackendConfig struct {
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight"`
}

type Config struct {
	Port     int             `yaml:"lb_port"`
	Backends []BackendConfig `yaml:"backends"`
}

func GetLBConfig() (*Config, error) {
	var config Config
	configfile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configfile, &config)
	if err != nil {
		return nil, err
	}
	if len(config.Backends) == 0 {
		return nil, errors.New("one or more Backends expected, none provided")
	}
	if config.Port == 0 {
		return nil, errors.New("port not found")
	}
	return &config, nil
}
