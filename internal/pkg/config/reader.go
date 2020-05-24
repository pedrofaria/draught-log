package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Resources Resources `yaml:"resources"`
}

type Resources []Resource

type Formatter struct {
	Type            string `yaml:"type"`
	MessageField    string `yaml:"messageField"`
	LevelField      string `yaml:"levelField"`
	TimestampField  string `yaml:"timestampField"`
	TimestampFormat string `yaml:"timestampFormat"`
}

type Resource struct {
	Name       string    `yaml:"name"`
	Provider   string    `yaml:"provider"`
	ProviderID string    `yaml:"providerId"`
	Formatter  Formatter `yaml:"formatter"`
}

func LoadFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
