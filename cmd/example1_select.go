package cmd

import (
	"aws/internal/example1"

	"github.com/spf13/cobra"
)

var example1SelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select tables with DynamoDB and Golang",
	Run:   example1.Select,
}

func init() {
	example1SelectCmd.Flags().StringP("tableName", "t", "test", "Table name")
}
