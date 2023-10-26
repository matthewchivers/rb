package config

import (
	"os"

	"github.com/matthewchivers/rb/pkg/fsutil"
	"github.com/matthewchivers/rb/pkg/pathparser"
	"gopkg.in/yaml.v2"
)

// Config is the struct representation of the config file
type Config struct {
	Rules []Rule `yaml:"rules"`
}

// Rule is the struct representation of a rule in the config file
type Rule struct {
	Name      string    `yaml:"name"`
	Directory string    `yaml:"directory"`
	Default   bool      `yaml:"default"`
	Match     MatchRule `yaml:"match,omitempty"`
}

// MatchRule is the struct representation of the match rule in the config file
type MatchRule struct {
	Repository RepositoryRule `yaml:"repository"`
}

// RepositoryRule is the struct representation of the repository rule in the config file
type RepositoryRule struct {
	Owner string `yaml:"owner,omitempty"`
	Name  string `yaml:"name,omitempty"`
}

// LoadConfig loads the config file from the specified path
func LoadConfig(filePath string) (*Config, error) {
	osfs := fsutil.OSFileSystem{}
	parsedPath, err := pathparser.ParsePath(osfs, filePath)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(parsedPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(parsedPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
