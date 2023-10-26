package mocks

import (
	"errors"
	"strings"
)

// MockFSUtil is a mock implementation of the FSUtil interface
type MockFSUtil struct{}

// UserHomeDir returns a mock user home directory
func (m *MockFSUtil) UserHomeDir() (string, error) {
	return "/home/user", nil
}

// Abs returns a mock absolute path
func (m *MockFSUtil) Abs(path string) (string, error) {
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
