package pathparser

import (
	"fmt"
	"strings"

	"github.com/matthewchivers/rb/pkg/fsutil"
)

// ParsePath checks for validity and expands a path
func ParsePath(fs fsutil.FSUtil, path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		return expandTilde(fs, path)
	}
	if strings.HasPrefix(path, ".") {
		return expandRelativePath(fs, path)
	}
	return path, nil
}

func expandTilde(fs fsutil.FSUtil, path string) (string, error) {
	home, err := fs.UserHomeDir()
	if err != nil {
		return path, fmt.Errorf("could not expand tilde path: %v", err)
	}
	return strings.Replace(path, "~", home, 1), nil
}

func expandRelativePath(fs fsutil.FSUtil, path string) (string, error) {
	absPath, err := fs.Abs(path)
	if err != nil {
		return path, fmt.Errorf("could not expand relative path: %v", err)
	}
	return absPath, nil
}
