package pathparser

import (
	"testing"

	"github.com/matthewchivers/rb/pkg/mocks"
)

func TestParsePath(t *testing.T) {
	fs := &mocks.MockFSUtil{}
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "expands tilde path",
			path:     "~/Documents",
			expected: "/home/user/Documents",
		},
		{
			name:     "expands relative current path",
			path:     "./config.yml",
			expected: "/current/config.yml",
		},
		{
			name:     "expands relative parent path",
			path:     "../config.yml",
			expected: "/parent/config.yml",
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
			actual, err := ParsePath(fs, tt.path)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tt.expected {
				t.Errorf("expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}
