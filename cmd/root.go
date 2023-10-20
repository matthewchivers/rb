package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type RBCommander struct {
	verbose bool
}

var commander = &RBCommander{}

var RootCmd = &cobra.Command{
	Use:   "rb [command] [arguments] [flags]",
	Short: "Repo Butler (rb) is a cli tool for managing git repositories.",
	Long:  "Repo Butler (rb) allows to to easily clone and access git repositories easily and efficiently.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("RootCmd PersistentPreRun")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&commander.verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
