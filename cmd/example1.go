package cmd

import (
	"github.com/spf13/cobra"
)

var example1Cmd = &cobra.Command{
	Use:   "example1",
	Short: "Simple examples of DynamoDB",
}

func init() {
	example1Cmd.AddCommand(example1ListCmd)
	example1Cmd.AddCommand(example1PutCmd)
}
