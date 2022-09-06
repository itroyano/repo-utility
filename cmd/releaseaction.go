package cmd

import (
	"github.com/itroyano/repo-utility/internal/actions"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for codecovCmd
var releaseActionCmd = &cobra.Command{
	Use: "release",
	// TODO: provide Short description for actions command
	Short: "Generate release build GitHub Action for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'release' command generates a release build GitHub Action for the repo.

Usage:
repo-utility actions release [flags]

Example:
repo-utility actions release
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		actions.GenerateRelease()
	},
}

func init() {
	actionsCmd.AddCommand(releaseActionCmd)
}
