package handlers

import (
	"fmt"

	gitutils "github.com/matthewchivers/rb/gitcore"
)

// Clone checks rules and clones a repository into a specified directory
func Clone(repositoryURL, customPath, rule string) {
	if rule != "" {
		// TODO: Implement rules - this is a placeholder for now to show how we can use the rule flag
		fmt.Printf("TODO: Implement rule: %s\n", rule)
	}

	err := gitutils.CloneGitRepo(repositoryURL, customPath)
	if err != nil {
		fmt.Println("Error cloning repository:", err)
	}

	fmt.Println("Repository cloned successfully!")
}
