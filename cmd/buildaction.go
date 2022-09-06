package cmd

import (
	"github.com/itroyano/repo-utility/internal/actions"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for codecovCmd
var buildActionCmd = &cobra.Command{
	Use: "main",
	// TODO: provide Short description for actions command
	Short: "Generate main build GitHub Action for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'main' command generates a main build GitHub Action for the repo.

Usage:
repo-utility actions main [flags]

Example:
repo-utility actions main
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		actions.GenerateMain()
	},
}

func init() {
	actionsCmd.AddCommand(buildActionCmd)
}
