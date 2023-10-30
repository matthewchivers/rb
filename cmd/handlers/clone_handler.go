package handlers

import (
	"fmt"

	gitutils "github.com/matthewchivers/rb/gitcore"
	"github.com/matthewchivers/rb/gitcore/utils/giturl"
	conf "github.com/matthewchivers/rb/pkg/config"
)

// Clone checks rules and clones a repository into a specified directory
func Clone(repositoryURL, customPath, customRule string) {
	cloneTargetPath := ""

	repo, err := giturl.Parse(repositoryURL)
	if err != nil {
		fmt.Println("Error parsing repository URL:", err)
		return
	}

	config, err := conf.LoadConfig("~/.config/rb/config.yaml")
	if err != nil {
		fmt.Println("Could not load config file:", err)
		if customPath == "" {
			fmt.Println("Please specify a path to clone the repository to")
			return
		}
		cloneTargetPath = customPath
		fmt.Println("Cloning into custom path:", cloneTargetPath)
	} else {
		cloneTargetPath, err = config.GetDirectory(repo)
		fmt.Println("Cloning into path:", cloneTargetPath)
		if err != nil {
			if _, ok := err.(*conf.NoRuleError); ok {
				fmt.Printf("%s", err)
			} else {
				fmt.Println("Error:", err)
				return
			}
		}
	}

	err = gitutils.CloneGitRepo(repositoryURL, cloneTargetPath)
	if err != nil {
		fmt.Println("Error cloning repository:", err)
	} else {
		fmt.Printf("Cloned repository: %s\nInto directory: %s\n", repositoryURL, cloneTargetPath)
	}
}
