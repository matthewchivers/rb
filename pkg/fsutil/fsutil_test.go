package fsutil

import (
	"errors"
	"testing"

	"github.com/matthewchivers/rb/pkg/mocks"
)

func TestFSUtil_UserHomeDir(t *testing.T) {
	fs := &mocks.MockFSUtil{}
	expected := "/home/user"
	actual, err := fs.UserHomeDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("expected %q, but got %q", expected, actual)
	}
}

func TestFSUtil_Abs(t *testing.T) {
	fs := &mocks.MockFSUtil{}
	tests := []struct {
		name     string
		path     string
		expected string
		err      error
	}{
		{
			name:     "valid current path",
			path:     "./valid/path",
			expected: "/current/valid/path",
			err:      nil,
		},
		{
			name:     "valid parent path",
			path:     "../valid/path",
			expected: "/parent/valid/path",
		},
		{
			name:     "valid absolute path",
			path:     "/valid/path",
			expected: "/valid/path",
		},
		{
			name:     "invalid path",
			path:     "invalid/path",
			expected: "",
			err:      errors.New("invalid path"),
		},
		{
			name:     "invalid empty path",
			path:     "",
			expected: "",
			err:      errors.New("invalid path"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := fs.Abs(tt.path)
			if err != nil && err.Error() != tt.err.Error() {
				t.Fatalf("expected error %v, but got %v", tt.err, err)
			}
			if actual != tt.expected {
				t.Errorf("expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}
