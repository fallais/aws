package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(example1Cmd)
	rootCmd.AddCommand(example2Cmd)
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
