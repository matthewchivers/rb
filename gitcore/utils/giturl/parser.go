package giturl

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/matthewchivers/rb/gitcore/types"
)

// Parse parses a git URL and returns a Repo struct
func Parse(repoURL string) (*types.Repo, error) {
	if strings.HasPrefix(repoURL, "git@") {
		return parseSSH(repoURL)
	}
	return parseHTTPS(repoURL)
}

func parseSSH(gitURL string) (*types.Repo, error) {
	re := regexp.MustCompile(`^git@(.*?):(.*?)/(.*?).git$`)
	matches := re.FindStringSubmatch(gitURL)

	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid git SSH URL format")
	}

	repo := &types.Repo{
		Host:  matches[1],
		Owner: matches[2],
		Name:  matches[3],
	}

	if !repo.Validate() {
		return nil, fmt.Errorf("invalid git SSH URL: does not meet validation criteria")
	}

	return repo, nil
}

func parseHTTPS(gitURL string) (*types.Repo, error) {
	url, err := url.Parse(gitURL)
	if err != nil {
		return nil, err
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("invalid git HTTP URL: scheme must be https")
	}
	parts := strings.Split(strings.Trim(url.Path, "/"), "/")
	if len(parts) != 2 || !strings.HasSuffix(parts[1], ".git") {
		return nil, fmt.Errorf("invalid git HTTP URL")
	}
	repo := &types.Repo{
		Host:  url.Host,
		Owner: parts[0],
		Name:  strings.TrimSuffix(parts[1], ".git"),
	}
	if !repo.Validate() {
		return nil, fmt.Errorf("invalid git HTTP URL: does not meet validation criteria")
	}
	return repo, nil
}
