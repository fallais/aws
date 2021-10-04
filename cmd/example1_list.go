package cmd

import (
	"aws/internal/example1"

	"github.com/spf13/cobra"
)

var example1ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tables with DynamoDB and Golang",
	Run:   example1.List,
}
