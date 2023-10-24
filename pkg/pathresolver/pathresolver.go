package pathresolver

import (
	"strings"

	"github.com/matthewchivers/rb/pkg/fsutil"
)

// ExpandPath expands a path to an absolute path
func ExpandPath(fs fsutil.FSUtil, path string) (string, error) {
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
		return path, err
	}
	return strings.Replace(path, "~", home, 1), nil
}

func expandRelativePath(fs fsutil.FSUtil, path string) (string, error) {
	absPath, err := fs.Abs(path)
	if err != nil {
		return path, err
	}
	return absPath, nil
}
