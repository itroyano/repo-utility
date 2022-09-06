package cmd

import (
	"github.com/itroyano/repo-utility/internal/codecov"
	"github.com/spf13/cobra"
)

// TODO: provide godoc compatible comment for codecovCmd
var codecovCmd = &cobra.Command{
	Use: "codecov",
	// TODO: provide Short description for actions command
	Short: "Generate CodeCov integration for the repo.",
	// TODO: provide Long description for actions command
	Long: `The 'codecov' command generates CodeCov integration for the repo.

Usage:
repo-utility codecov [flags]

Example:
repo-utility codecov
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		codecov.GenerateCodecov()
	},
}

func init() {
	rootCmd.AddCommand(codecovCmd)
}
