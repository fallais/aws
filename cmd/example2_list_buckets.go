package cmd

import (
	"aws/internal/example2"

	"github.com/spf13/cobra"
)

var example2ListBucketsCmd = &cobra.Command{
	Use:   "list_buckets",
	Short: "List buckets with S3 and Golang",
	Run:   example2.ListBuckets,
}
