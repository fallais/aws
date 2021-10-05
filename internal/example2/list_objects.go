package example2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// ListObjects is a convenient function for Cobra.
func ListObjects(cmd *cobra.Command, args []string) {
	// Get the bucket name from flags
	bucketName, err := cmd.Flags().GetString("bucketName")
	if err != nil {
		log.Fatal("error with bucketName flag")
	}

	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the S3 client
	client := s3.NewFromConfig(cfg)

	// Set the parameters
	params := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	}

	// List tables (they recommand using V2)
	resp, err := client.ListObjectsV2(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	if len(resp.Contents) == 0 {
		log.Print("no object has been found")
	} else {
		log.Printf("%d objects has been found", len(resp.Contents))
	}

	// Process the buckets
	for _, bucket := range resp.Contents {
		// Print information
		log.Print(bucket.Key, bucket.Size)
	}

	log.Printf("Successfully listed objects for bucket [%s]", bucketName)
}
