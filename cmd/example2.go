package cmd

import (
	"github.com/spf13/cobra"
)

var example2Cmd = &cobra.Command{
	Use:   "example2",
	Short: "Simple examples of using S3 with Golang",
}

func init() {
	//example2Cmd.AddCommand(example2ListCmd)
	//example2Cmd.AddCommand(example2PutCmd)
}
