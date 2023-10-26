package config

import (
	"fmt"
	"regexp"

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

func getRulePath(rule Rule, repo *types.Repo) (string, error) {
	rulePath, err := pathparser.ParsePath(fsutil.OSFileSystem{}, rule.Directory)
	if err != nil {
		return "", err
	}
	clonePath := rulePath + "/" + repo.Owner + "/" + repo.Name
	return clonePath, nil
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
