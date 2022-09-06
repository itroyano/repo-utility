package cmd

import (
	"github.com/itroyano/repo-utility/internal/coveralls"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for coverallsCmd
var coverallsCmd = &cobra.Command{
	Use: "coveralls",
	// TODO: provide Short description for actions command
	Short: "Generate CoverAlls integration for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'coveralls' command generates CoverAlls integration for the repo.

Usage:
repo-utility coveralls [flags]

Example:
repo-utility coveralls
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		coveralls.GenerateCoveralls()
	},
}

func init() {
	rootCmd.AddCommand(coverallsCmd)
}
