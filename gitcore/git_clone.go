package gitutils

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/matthewchivers/rb/gitcore/utils/giturl"
)

// CloneGitRepo clones a git repository into a specified directory
func CloneGitRepo(repoURL, destDir string) error {
	if _, err := giturl.Parse(repoURL); err != nil {
		return err
	}

	if err := ensureDirExists(destDir); err != nil {
		return err
	}

	if _, err := git.PlainClone(destDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	}); err != nil {
		return err
	}
	return nil
}

// ensureDirExists checks if a directory exists, and creates it if it doesn't
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
