package cmd

import (
	"fmt"

	"github.com/matthewchivers/rb/gitutils"
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

func CloneHandler(cmd *cobra.Command, args []string) {
	repositoryURL := args[0]
	if customPath == "" {
		fmt.Errorf("Error: No target path specified\n")
		return
	}
	cloneTargetPath := customPath // Simplify by directly using customPath
	if rule != "" {
		fmt.Printf("TODO: Implement rule: %s\n", rule)
	}
	if err := gitutils.CloneGitRepo(repositoryURL, cloneTargetPath); err != nil { // Simplified error handling
		fmt.Println("Error cloning repository:", err)
	}
}
