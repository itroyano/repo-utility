package cmd

import (
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for actionsCmd
var actionsCmd = &cobra.Command{
	Use: "actions",
	// TODO: provide Short description for actions command
	Short: "Generate GitHub Actions for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'actions' command generates GitHub Actions of various types, to integrate a CI test, build or release processes into the repository.

Usage:
repo-utility actions [flags]

Example:
repo-utility actions build
`,
}

func init() {
	rootCmd.AddCommand(actionsCmd)
}
