package utils

import (
	"os"
	"path/filepath"
)

// FileSystem is an interface for interacting with the filesystem
type FileSystem interface {
	UserHomeDir() (string, error)
	Abs(path string) (string, error)
}

// OSFileSystem implements FileSystem using the os and filepath packages
type OSFileSystem struct{}

// UserHomeDir returns the current user's home directory
func (OSFileSystem) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// Abs returns the absolute path of a given path
func (OSFileSystem) Abs(path string) (string, error) {
	return filepath.Abs(path)
}
