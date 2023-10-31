package pathparser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParsePath(t *testing.T) {
	home, _ := os.UserHomeDir()
	absPathCurrent, _ := filepath.Abs(".")
	absPathParent, _ := filepath.Abs("..")
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "expands tilde path",
			path:     "~/Documents",
			expected: fmt.Sprintf("%s/Documents", home),
		},
		{
			name:     "expands relative current path",
			path:     "./config.yml",
			expected: fmt.Sprintf("%s/config.yml", absPathCurrent),
		},
		{
			name:     "expands relative parent path",
			path:     "../config.yml",
			expected: fmt.Sprintf("%s/config.yml", absPathParent),
		},
		{
			name:     "returns path as is",
			path:     "/usr/local/bin",
			expected: "/usr/local/bin",
		},
		{
			name:     "returns normalised path (removes trailing slash)",
			path:     "/usr/local/bin/",
			expected: "/usr/local/bin",
		},
		{
			name:     "returns normalised path (adds leading slash)",
			path:     "usr/local/bin",
			expected: "/usr/local/bin",
		},
		{
			name:     "returns normalised path (adds leading slash and removes trailing slash)",
			path:     "usr/local/bin/",
			expected: "/usr/local/bin",
		},
		{
			name:     "returns normalised path (even with multiple slashes)",
			path:     "usr///local//bin////",
			expected: "/usr/local/bin",
		},
		{
			name:     "returns normalised path (even with multiple leding or trailing slashes)",
			path:     "///usr/local/bin////",
			expected: "/usr/local/bin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ParsePath(tt.path)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tt.expected {
				t.Errorf("expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}
