package cmd

import (
	"aws/internal/example1"

	"github.com/spf13/cobra"
)

var example1PutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put stuff into DynamoDB table with Golang",
	Run:   example1.Put,
}

func init() {
	example1PutCmd.Flags().StringP("tableName", "t", "test", "Table name")
}
