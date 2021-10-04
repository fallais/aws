package cmd

import (
	"aws/internal/example2"

	"github.com/spf13/cobra"
)

var example2ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List buckets with S3 and Golang",
	Run:   example2.List,
}
