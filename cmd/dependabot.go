package cmd

import (
	"github.com/itroyano/repo-utility/internal/dependabot"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for dependabotCmd
var dependabotCmd = &cobra.Command{
	Use: "actions",
	// TODO: provide Short description for actions command
	Short: "Generate Dependabot configuration for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'dependabot' command generates dependabot configuration for the repo.

Usage:
repo-utility dependabot [flags]

Example:
repo-utility dependabot
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dependabot.GenerateBot()
	},
}

func init() {
	rootCmd.AddCommand(dependabotCmd)
}
