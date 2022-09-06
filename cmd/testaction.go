package cmd

import (
	"github.com/itroyano/repo-utility/internal/actions"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for codecovCmd
var testActionCmd = &cobra.Command{
	Use: "test",
	// TODO: provide Short description for actions command
	Short: "Generate test CI build GitHub Action for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'test' command generates a test CI build GitHub Action for the repo.

Usage:
repo-utility actions test [flags]

Example:
repo-utility actions test
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		actions.GenerateTest()
	},
}

func init() {
	actionsCmd.AddCommand(testActionCmd)
}
