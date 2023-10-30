package config

import (
	"os"

	"github.com/matthewchivers/rb/pkg/fsutil"
	rblogger "github.com/matthewchivers/rb/pkg/logger"
	"github.com/matthewchivers/rb/pkg/pathparser"
	"gopkg.in/yaml.v2"
)

var (
	logger = rblogger.GetLogger()
)

// Directory represents a directory in the config file
type Directory struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// Repository represents a repository in the config file
type Repository struct {
	URL           string `yaml:"url"`
	DirectoryName string `yaml:"directory_name"`
	RelativePath  string `yaml:"relative_path"`
}

// MatchRepository represents a repository in a match rule
type MatchRepository struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
	Host  string `yaml:"host"`
}

// CloneRule represents a clone rule in the config file
type CloneRule struct {
	Name           string `yaml:"name"`
	Default        bool   `yaml:"default"`
	DirectoryName  string `yaml:"directory_name"`
	NestingPattern string `yaml:"nesting_pattern"`
	Match          struct {
		// TODO: Add support for multiple matches
		Repository MatchRepository `yaml:"repository"`
	} `yaml:"match"`
}

// Config represents the config file
type Config struct {
	Directories  []Directory  `yaml:"directories"`
	Repositories []Repository `yaml:"repositories"`
	CloneRules   []CloneRule  `yaml:"clone_rules"`
}

// LoadConfig loads the config file from the specified path
func LoadConfig(filePath string) (*Config, error) {
	logger.Infof("Loading config file from %s", filePath)
	osfs := fsutil.OSFileSystem{}
	parsedPath, err := pathparser.ParsePath(osfs, filePath)
	if err != nil {
		logger.Errorf("Could not parse path: %s", err)
		return nil, err
	}

	_, err = os.Stat(parsedPath)
	if err != nil {
		logger.Errorf("File does not exist: %s", err)
		return nil, err
	}

	logger.Debugf("Reading config file %s", parsedPath)
	data, err := os.ReadFile(parsedPath)
	if err != nil {
		return nil, err
	}
	logger.Debugf("Config file contents: %s", data)

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	logger.Infof("Loaded config file from %s", filePath)
	return &config, nil
}
