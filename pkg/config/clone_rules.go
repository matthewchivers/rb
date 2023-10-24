package config

import (
	"regexp"

	"github.com/matthewchivers/rb/pkg/fsutil"
	pr "github.com/matthewchivers/rb/pkg/pathresolver"
)

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
	return pr.ExpandPath(fsutil.OSFileSystem{}, defaultRule.Directory)
}
