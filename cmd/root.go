package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:           "mtemp",
	Short:         "Create and monitor temp emails",
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {
	return rootCmd.Execute()
}
