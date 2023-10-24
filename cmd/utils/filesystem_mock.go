package utils

import (
	"errors"
	"strings"
)

type mockFileSystem struct{}

func (m *mockFileSystem) UserHomeDir() (string, error) {
	return "/home/user", nil
}

func (m *mockFileSystem) Abs(path string) (string, error) {
	if path == "" {
		return "", errors.New("invalid path")
	}

	if strings.HasPrefix(path, "/") {
		// Generic valid path (no resolution needed)
		return path, nil
	}

	// Handle cases where path starts with "./" or "../"
	if strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") {
		resolvedPath := path
		for strings.HasPrefix(resolvedPath, "./") {
			resolvedPath = strings.Replace(resolvedPath, "./", "/current/", 1)
		}
		for strings.HasPrefix(resolvedPath, "../") {
			resolvedPath = strings.Replace(resolvedPath, "../", "/parent/", 1)
		}
		return resolvedPath, nil
	}

	return "", errors.New("invalid path")
}
