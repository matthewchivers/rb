package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/matthewchivers/rb/gitcore/types"
	"github.com/matthewchivers/rb/pkg/fsutil"
	"github.com/matthewchivers/rb/pkg/pathparser"
)

// NoRuleError is returned when no rule matches the repository
type NoRuleError struct {
	Msg  string
	Code int
}

func (e *NoRuleError) Error() string {
	return fmt.Sprintf("Error Code %d: %s", e.Code, e.Msg)
}

// GetDirectory returns the directory to clone the repository into
// based on the rules in the config file.
// If no rules match, an error is returned.
func (c *Config) GetDirectory(repo *types.Repo) (string, error) {
	logger.Infof("Getting directory for repository %s", repo.Name)
	var defaultRule *CloneRule
	for _, rule := range c.CloneRules {
		logger.Debugf("Checking rule %s", rule.Name)
		if checkRepoMatches(rule, repo) {
			logger.Debugf("Rule %s matched the repository", rule.Name)
			return c.getRulePath(rule, repo)
		}
		if rule.Default {
			logger.Debugf("Rule %s is default", rule.Name)
			defaultRule = &rule
		}
	}
	if defaultRule != nil {
		logger.Debugf("No rule matched the repository, using default rule %s", defaultRule.Name)
		return c.getRulePath(*defaultRule, repo)
	}
	logger.Debugf("No rule matched the repository")
	return "", &NoRuleError{Msg: "No rule matched the repository", Code: 1}
}

// getRulePath returns the path to clone the repository into based on the rule.
func (c *Config) getRulePath(rule CloneRule, repo *types.Repo) (string, error) {
	dirMap := make(map[string]string)
	for _, directory := range c.Directories {
		dirMap[directory.Name] = directory.Path
	}
	logger.Debugf("Directories: %v", dirMap)
	if _, ok := dirMap[rule.DirectoryName]; !ok {
		return "", fmt.Errorf("directory name %s not found in config", rule.DirectoryName)
	}
	logger.Infof("Directory name %s found in rule %s", rule.DirectoryName, rule.Name)
	path := dirMap[rule.DirectoryName]
	logger.Debugf("Base path: %s", path)
	if rule.NestingPattern != "" {
		nestedPath, err := getNestedPath(rule.NestingPattern, repo)
		logger.Debugf("Nested path: %s", nestedPath)
		if err != nil {
			return "", err
		}
		// return dirMap[rule.DirectoryName] + nestedPath, nil
		path = path + nestedPath
	}
	repoPath := path + "/" + repo.Name
	logger.Debugf("Rule path: %s", repoPath)
	return repoPath, nil
}

// getNestedPath processes the nesting pattern and returns the nested path.
func getNestedPath(pattern string, repo *types.Repo) (string, error) {
	// Validate the pattern
	logger.Debugf("Validating pattern: %s", pattern)
	validPattern := regexp.MustCompile(`^([\w/]*\{(?:host|owner)\}[\w/]*)+$`)
	if !validPattern.MatchString(pattern) {
		return "", fmt.Errorf("invalid pattern: %s", pattern)
	}
	logger.Debugf("Pattern %s is valid", pattern)
	// Perform replacements
	nestedPath := strings.ReplaceAll(pattern, "{host}", repo.Host)
	nestedPath = strings.ReplaceAll(nestedPath, "{owner}", repo.Owner)
	nestedPath, err := pathparser.ParsePath(fsutil.OSFileSystem{}, nestedPath)
	if err != nil {
		return "", err
	}
	logger.Infof("Nested path: %s", nestedPath)
	return nestedPath, nil
}

// checkRepoMatches checks if the repository matches the rule.
func checkRepoMatches(rule CloneRule, repo *types.Repo) bool {
	allMatched := true
	logger.Debugf("Checking if repository %s matches rule %s", repo.Name, rule.Name)
	if rule.Match.Repository.Owner != "" {
		matched, _ := regexp.MatchString(rule.Match.Repository.Owner, repo.Owner)
		allMatched = allMatched && matched
	}
	if rule.Match.Repository.Name != "" {
		matched, _ := regexp.MatchString(rule.Match.Repository.Name, repo.Name)
		allMatched = allMatched && matched
	}
	if rule.Match.Repository.Host != "" {
		matched, _ := regexp.MatchString(rule.Match.Repository.Host, repo.Host)
		allMatched = allMatched && matched
	}
	return allMatched
}
