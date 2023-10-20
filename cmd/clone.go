package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	customPath string
	rule       string
)

var cloneCmd = &cobra.Command{
	Use:   "clone [repo-url]",
	Short: "clone a repository into a directory",
	Long:  "clone a repository into a directory based on the user-specified ruleset",
	Args:  cobra.ExactArgs(1),
	Run:   CloneHandler,
}

func init() {
	cloneCmd.Flags().StringVarP(&customPath, "path", "p", "", "custom path to clone repository into")
	cloneCmd.Flags().StringVarP(&rule, "rule", "r", "", "rule to apply when cloning repository")
	RootCmd.AddCommand(cloneCmd)
}

func CloneHandler(cmd *cobra.Command, args []string) {
	repositoryURL := args[0]
	if customPath != "" {
		fmt.Printf("Cloning repository: %s into %s\n", repositoryURL, customPath)
	} else if rule != "" {
		fmt.Printf("Cloning repository: %s using rule: %s\n", repositoryURL, rule)
	} else {
		fmt.Printf("Cloning repository: %s\n", repositoryURL)
	}
}
