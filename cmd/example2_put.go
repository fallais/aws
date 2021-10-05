package cmd

import (
	"aws/internal/example2"

	"github.com/spf13/cobra"
)

var example2PutCmd = &cobra.Command{
	Use:   "example2",
	Short: "Simple examples of using S3 with Golang",
	Run:   example2.Put,
}

func init() {
	example1PutCmd.Flags().StringP("bucketName", "n", "test", "Bucket name")
}
