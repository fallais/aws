package cmd

import (
	"github.com/spf13/cobra"
)

var example2Cmd = &cobra.Command{
	Use:   "example2",
	Short: "Simple examples of using S3 with Golang",
}

func init() {
	example2Cmd.AddCommand(example2ListBucketsCmd)
	example2Cmd.AddCommand(example2ListObjectsCmd)
	example2Cmd.AddCommand(example2PutCmd)
}
