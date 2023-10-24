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
	parts := strings.Split(repoURL, ":")
	if len(parts) != 2 {
		return false
	} else if parts[0] != user+"@"+host {
		return false
	}
	if !strings.HasSuffix(parts[1], ".git") {
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
	if !strings.HasSuffix(parsedURL.Path, ".git") {
		return false
	}
	return true
}
