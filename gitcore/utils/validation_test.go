package utils

import "testing"

func TestCheckGitHURL(t *testing.T) {
	tests := []struct {
		name    string
		repoURL string
		user    string
		host    string
		want    bool
	}{
		{
			name:    "valid ssh url",
			repoURL: "git@github.com:owner/repository.git",
			want:    true,
		},
		{
			name:    "valid https url",
			repoURL: "https://github.com/owner/repository.git",
			want:    true,
		},
		{
			name:    "invalid ssh url - missing .git suffix",
			repoURL: "git@github.com:owner/repository",
			want:    false,
		},
		{
			name:    "invalid ssh url - missing git@ prefix",
			repoURL: "github.com:owner/repository.git",
			want:    false,
		},
		{
			name:    "invalid ssh url - missing host",
			repoURL: "git@:owner/repository.git",
			want:    false,
		},
		{
			name:    "invalid ssh url - missing owner",
			repoURL: "git@github.com:/repository.git",
			want:    false,
		},
		{
			name:    "invalid ssh url - missing owner and host",
			repoURL: "git@:/repository.git",
			want:    false,
		},
		{
			name:    "invalid https url - missing .git suffix",
			repoURL: "https://github.com/owner/repository",
			want:    false,
		},
		{
			name:    "invalid https url - missing https:// prefix",
			repoURL: "github.com/owner/repository.git",
			want:    false,
		},
		{
			name:    "invalid https url - missing host",
			repoURL: "https:///owner/repository.git",
			want:    false,
		},
		{
			name:    "invalid https url - missing owner",
			repoURL: "https://github.com//repository.git",
		},
		{
			name:    "invalid https url - missing owner and host",
			repoURL: "https:///repository.git",
			want:    false,
		},
		{
			name:    "invalid url - missing scheme",
			repoURL: "github.com/owner/repository.git",
			want:    false,
		},
		{
			name:    "invalid url - empty",
			repoURL: "",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidGitURL(tt.repoURL); got != tt.want {
				t.Errorf("checkSSHURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
