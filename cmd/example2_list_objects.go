package cmd

import (
	"aws/internal/example2"

	"github.com/spf13/cobra"
)

var example2ListObjectsCmd = &cobra.Command{
	Use:   "list_objects",
	Short: "List bucket objects with S3 and Golang",
	Run:   example2.ListObjects,
}
