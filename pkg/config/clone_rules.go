package config

import (
	"errors"
	"regexp"

	"github.com/matthewchivers/rb/pkg/fsutil"
	pr "github.com/matthewchivers/rb/pkg/pathresolver"
)

// GetDirectory returns the directory to clone the repository into
// based on the rules in the config file.
// If no rules match, an error is returned.
func (c *Config) GetDirectory(repoOwner, repoName string) (string, error) {
	defaultRule := Rule{}
	for _, rule := range c.Rules {
		if rule.Match.Repository.Owner != "" {
			matched, _ := regexp.MatchString(rule.Match.Repository.Owner, repoOwner)
			if matched {
				return pr.ExpandPath(fsutil.OSFileSystem{}, rule.Directory)
			}
		}
		if rule.Match.Repository.Name != "" {
			matched, _ := regexp.MatchString(rule.Match.Repository.Name, repoName)
			if matched {
				return pr.ExpandPath(fsutil.OSFileSystem{}, rule.Directory)
			}
		}
		if rule.Default {
			defaultRule = rule
		}
	}
	if defaultRule.Directory == "" {
		return "", errors.New("no matching rule found")
	}
	return pr.ExpandPath(fsutil.OSFileSystem{}, defaultRule.Directory)
}
