package utils

import (
	"net/url"
	"strings"
)

// IsValidGitURL checks if a git URL is valid
func IsValidGitURL(repoURL string) bool {
	// TODO: Add support for other git providers
	return isValidGitHubURL(repoURL)
}

func isValidGitHubURL(repoURL string) bool {
	if strings.HasPrefix(repoURL, "git@") {
		return checkSSHURL(repoURL, "git", "github.com")
	}
	return checkHTTPSURL(repoURL, "github.com")
}

func checkSSHURL(repoURL, user, host string) bool {
	// git@<host>:<owner>/<repo>.git
	colonSplitParts := strings.Split(repoURL, ":")
	if len(colonSplitParts) != 2 {
		return false
	}
	if colonSplitParts[0] != user+"@"+host {
		return false
	}
	slashSplitParts := strings.Split(colonSplitParts[1], "/")
	if len(slashSplitParts) != 2 {
		return false
	}
	if slashSplitParts[0] == "" || slashSplitParts[1] == "" {
		return false
	}
	if !strings.HasSuffix(colonSplitParts[1], ".git") {
		return false
	}
	return true
}

func checkHTTPSURL(repoURL, host string) bool {
	// https://<host>/<owner>/<repo>.git
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return false
	}
	if parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host != host {
		return false
	}
	slashSplitParts := strings.Split(parsedURL.Path, "/")
	if len(slashSplitParts) != 3 {
		return false
	}
	if slashSplitParts[1] == "" || slashSplitParts[2] == "" {
		return false
	}
	if !strings.HasSuffix(parsedURL.Path, ".git") {
		return false
	}
	return true
}
