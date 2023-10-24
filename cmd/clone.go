package cmd

import (
	"errors"
	"fmt"

	"github.com/matthewchivers/rb/cmd/handlers"
	"github.com/matthewchivers/rb/pkg/fsutil"
	pr "github.com/matthewchivers/rb/pkg/pathresolver"
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
	Run: func(cmd *cobra.Command, args []string) {
		repositoryURL := args[0]
		cloneTargetPath, err := checkCustomPath(customPath)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		handlers.Clone(repositoryURL, cloneTargetPath, rule)
	},
}

func init() {
	cloneCmd.Flags().StringVarP(&customPath, "path", "p", "", "Custom path to clone repository into")
	cloneCmd.Flags().StringVarP(&rule, "rule", "r", "", "Rule to apply when cloning repository")
	RootCmd.AddCommand(cloneCmd)
}

// checkCustomPath checks if a custom path has been specified. If so, returns the expanded path
func checkCustomPath(customPath string) (string, error) {
	if customPath == "" {
		return "", errors.New("custom path must be specified")
	}
	osfs := fsutil.OSFileSystem{}
	cloneTargetPath, err := pr.ExpandPath(osfs, customPath)
	if err != nil {
		return "", fmt.Errorf("error expanding path: %w", err)
	}
	return cloneTargetPath, nil
}
