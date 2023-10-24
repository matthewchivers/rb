package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Name      string    `yaml:"name"`
	Directory string    `yaml:"directory"`
	Default   bool      `yaml:"default"`
	Match     MatchRule `yaml:"match,omitempty"`
}

type MatchRule struct {
	Repository RepositoryRule `yaml:"repository"`
}

type RepositoryRule struct {
	Owner string `yaml:"owner,omitempty"`
	Name  string `yaml:"name,omitempty"`
}

func LoadConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
