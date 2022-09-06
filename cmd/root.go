package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "repo-utility",
	Short: "Repo Utility generates commonly used CI integrations",
	Long: `Repo Utility generates commonly used CI integrations in a code repository.
	These include GitHub actions, dependency management, code coverage, and more`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "repo-utility tool execution failed: ", err)
		os.Exit(1)
	}
}
