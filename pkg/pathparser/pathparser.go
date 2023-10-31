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
	parsedPath := normalisePath(path)
	var err error
	switch {
	case strings.HasPrefix(path, "~"):
		parsedPath, err = expandTilde(path)
	case strings.HasPrefix(path, "."):
		parsedPath, err = expandRelativePath(path)
	}
	if err != nil {
		return path, err
	}
	return parsedPath, nil
}

func normalisePath(path string) string {
	// Replace multiple consecutive slashes with a single slash
	normalisedPath := regexp.MustCompile(`//+`).ReplaceAllString(path, "/")
	// Ensure prefix is a slash
	if !strings.HasPrefix(path, "/") {
		normalisedPath = "/" + normalisedPath
	}
	// Remove trailing slash
	return strings.TrimSuffix(normalisedPath, "/")
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
