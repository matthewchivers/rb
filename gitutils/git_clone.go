package gitutils

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

// CloneGitRepo clones a git repository into a specified directory
func CloneGitRepo(repoURL, destDir string) error {
	if !checkValidGitURL(repoURL) {
		return fmt.Errorf("invalid git URL: %s", repoURL)
	}
	destDir = expandHomeDir(destDir)

	if err := ensureDirExists(destDir); err != nil {
		return err
	}

	_, err := git.PlainClone(destDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return err
}

func checkValidGitURL(repoURL string) bool {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return false
	}
	if parsedURL.Scheme != "git" &&
		parsedURL.Scheme != "https" &&
		parsedURL.Scheme != "ssh" &&
		parsedURL.Scheme != "http" {
		return false
	}

	if !strings.HasSuffix(parsedURL.Path, ".git") {
		return false
	}

	return true
}

// expandHomeDir expands the ~ to the user's home directory
func expandHomeDir(path string) string {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return path // return original path if home directory can't be determined
		}
		return filepath.Join(homeDir, path[2:])
	}
	return path
}

// ensureDirExists checks if a directory exists, and creates it if it doesn't
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
