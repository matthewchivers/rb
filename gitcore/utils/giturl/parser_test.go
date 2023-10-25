package giturl

import (
	"testing"

	"github.com/matthewchivers/rb/gitcore/types"
)

func TestCheckGitHURL(t *testing.T) {
	tests := []struct {
		name         string
		repoURL      string
		expectedRepo *types.Repo
		expectedErr  bool
	}{
		{
			name:         "valid ssh url",
			repoURL:      "git@github.com:owner/repository.git",
			expectedRepo: &types.Repo{Owner: "owner", Name: "repository", Host: "github.com"},
		},
		{
			name:         "valid https url",
			repoURL:      "https://github.com/owner/repository.git",
			expectedRepo: &types.Repo{Owner: "owner", Name: "repository", Host: "github.com"},
		},
		{
			name:         "invalid ssh url - missing .git suffix",
			repoURL:      "git@github.com:owner/repository",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid ssh url - missing git@ prefix",
			repoURL:      "github.com:owner/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid ssh url - missing host",
			repoURL:      "git@:owner/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid ssh url - missing owner",
			repoURL:      "git@github.com:/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid ssh url - missing owner and host",
			repoURL:      "git@:/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid https url - missing .git suffix",
			repoURL:      "https://github.com/owner/repository",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid https url - missing https:// prefix",
			repoURL:      "github.com/owner/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid https url - missing host",
			repoURL:      "https:///owner/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid https url - missing owner",
			repoURL:      "https://github.com//repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid https url - missing owner and host",
			repoURL:      "https:///repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid url - missing scheme",
			repoURL:      "github.com/owner/repository.git",
			expectedRepo: nil,
			expectedErr:  true,
		},
		{
			name:         "invalid url - empty",
			repoURL:      "",
			expectedRepo: nil,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRepo, err := Parse(tt.repoURL)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Parse() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if actualRepo != nil {
				if *actualRepo != *tt.expectedRepo {
					t.Errorf("Parse() actualRepo = %v, expected %v", *actualRepo, *tt.expectedRepo)
				}
			}
		})
	}
}
