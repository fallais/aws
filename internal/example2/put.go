package example2

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// Put is a convenient function for Cobra.
func Put(cmd *cobra.Command, args []string) {
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
	params := &s3.PutObjectInput{
		Bucket: &bucketName,
		Body:   strings.NewReader("fake file"),
	}

	// Put object into bucket
	_, err = client.PutObject(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to put object into bucket [%s]: %v", bucketName, err)
	}

	log.Printf("Successfully put object into bucket [%s]", bucketName)
}
