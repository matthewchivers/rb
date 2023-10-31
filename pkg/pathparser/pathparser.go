package pathparser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParsePath checks for validity and expands a path
func ParsePath(path string) (string, error) {
	parsedPath := path
	if strings.HasPrefix(path, "~") {
		tildePath, err := expandTilde(parsedPath)
		if err != nil {
			return path, err
		}
		parsedPath = tildePath
	} else if strings.HasPrefix(path, ".") {
		dotPath, err := expandRelativePath(parsedPath)
		if err != nil {
			return path, err
		}
		parsedPath = dotPath
	}
	parsedPath = normalisePath(parsedPath)
	return parsedPath, nil
}

func normalisePath(path string) string {
	normalisedPath := path
	// Replace multiple consecutive slashes with a single slash
	normalisedPath = regexp.MustCompile(`//+`).ReplaceAllString(normalisedPath, "/")
	// Ensure prefix is a slash
	if !strings.HasPrefix(path, "/") {
		normalisedPath = "/" + normalisedPath
	}
	// Remove trailing slash
	normalisedPath = strings.TrimSuffix(normalisedPath, "/")
	return normalisedPath
}

func expandTilde(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return path, fmt.Errorf("could not expand tilde path: %v", err)
	}
	return strings.Replace(path, "~", home, 1), nil
}

func expandRelativePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path, fmt.Errorf("could not expand relative path: %v", err)
	}
	return absPath, nil
}
