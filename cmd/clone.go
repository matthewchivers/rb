package cmd

import (
	"fmt"

	gitutils "github.com/matthewchivers/rb/gitcore"
	"github.com/spf13/cobra"
)

var (
	customPath string
	rule       string
)

var cloneCmd = &cobra.Command{
	Use:   "clone [repo-url]",
	Short: "Clone a repository into a directory",
	Long:  "Clone a repository into a directory based on the user-specified ruleset",
	Args:  cobra.ExactArgs(1),
	Run:   CloneHandler,
}

func init() {
	cloneCmd.Flags().StringVarP(&customPath, "path", "p", "", "Custom path to clone repository into")
	cloneCmd.Flags().StringVarP(&rule, "rule", "r", "", "Rule to apply when cloning repository")
	RootCmd.AddCommand(cloneCmd)
}

// CloneHandler handles the clone command
func CloneHandler(cmd *cobra.Command, args []string) {
	repositoryURL := args[0]

	// Until we implement rules, we require a custom path to be specified
	cloneTargetPath := customPath
	if customPath == "" {
		fmt.Println("custom path must be specified")
		return
	}

	if rule != "" {
		// TODO: Implement rules - this is a placeholder for now to show how we can use the rule flag
		fmt.Printf("TODO: Implement rule: %s\n", rule)
	}
	if err := gitutils.CloneGitRepo(repositoryURL, cloneTargetPath); err != nil {
		fmt.Println("Error cloning repository:", err)
	}
	fmt.Println("Repository cloned successfully!")
}
