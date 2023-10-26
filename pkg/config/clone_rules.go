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
	defaultRule := Rule{}
	for _, rule := range c.Rules {
		allMatched := checkRepoMatches(rule, repo)
		if allMatched {
			return getRulePath(rule, repo)
		}
		if rule.Default {
			defaultRule = rule
		}
	}
	if defaultRule.Directory == "" {
		return "", &NoRuleError{
			Msg:  "No rule found for repository",
			Code: 1,
		}
	}
	return pathparser.ParsePath(fsutil.OSFileSystem{}, defaultRule.Directory)
}

// getRulePath returns the path to clone the repository into based on the supplied rule
func getRulePath(rule Rule, repo *types.Repo) (string, error) {
	rulePath, err := pathparser.ParsePath(fsutil.OSFileSystem{}, rule.Directory)
	if err != nil {
		return "", err
	}

	var clonePath string
	if rule.Nesting.Pattern != "" {
		nestedPath, err := getNestedPath(rule.Nesting.Pattern, repo)
		if err != nil {
			return "", err
		}
		clonePath = rulePath + nestedPath
	} else {
		clonePath = rulePath
	}

	return clonePath + "/" + repo.Name, nil
}

// getNestedPath processes the nesting pattern and returns the nested path.
func getNestedPath(pattern string, repo *types.Repo) (string, error) {
	// Validate the pattern
	validPattern := regexp.MustCompile(`^([\w/]*\{(?:host|owner)\}[\w/]*)+$`)
	if !validPattern.MatchString(pattern) {
		return "", fmt.Errorf("invalid pattern: %s", pattern)
	}
	// Perform replacements
	nestedPath := strings.ReplaceAll(pattern, "{host}", repo.Host)
	nestedPath = strings.ReplaceAll(nestedPath, "{owner}", repo.Owner)
	nestedPath, err := pathparser.ParsePath(fsutil.OSFileSystem{}, nestedPath)
	if err != nil {
		return "", err
	}
	return nestedPath, nil
}

func checkRepoMatches(rule Rule, repo *types.Repo) bool {
	allMatched := true
	if rule.Match.Repository.Owner != "" {
		matched, _ := regexp.MatchString(rule.Match.Repository.Owner, repo.Owner)
		allMatched = allMatched && matched
	}
	if rule.Match.Repository.Name != "" {
		matched, _ := regexp.MatchString(rule.Match.Repository.Name, repo.Name)
		allMatched = allMatched && matched
	}
	return allMatched
}
