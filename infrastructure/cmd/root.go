package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	counterCommand.Flags().StringVarP(&word, "word", "w", "", "Word or a pattern to search through a directory.")
	counterCommand.Flags().StringVarP(&directory, "dir", "d", "", "Where to search")
	counterCommand.MarkFlagRequired("word")
	counterCommand.MarkFlagRequired("dir")

	rootCmd.AddCommand(runserverCommand)
	rootCmd.AddCommand(counterCommand)
}
