package gitutils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/matthewchivers/rb/gitcore/utils"
)

// CloneGitRepo clones a git repository into a specified directory
func CloneGitRepo(repoURL, destDir string) error {
	if !utils.IsValidGitURL(repoURL) {
		return fmt.Errorf("invalid git URL: %s", repoURL)
	}

	if err := ensureDirExists(destDir); err != nil {
		return err
	}

	_, err := git.PlainClone(destDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return err
}

// ensureDirExists checks if a directory exists, and creates it if it doesn't
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
