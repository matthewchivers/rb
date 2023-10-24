package fsutil

import (
	"os"
	"path/filepath"
)

// FSUtil is an interface for interacting with the filesystem
type FSUtil interface {
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
