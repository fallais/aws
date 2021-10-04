package example2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// List is a convenient function for Cobra.
func List(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the S3 client
	client := s3.NewFromConfig(cfg)

	// List tables
	resp, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	if len(resp.Buckets) == 0 {
		log.Print("no table has been found")
	} else {
		log.Printf("%d tables has been found", len(resp.Buckets))
	}

	// Process the buckets
	for _, bucket := range resp.Buckets {
		// Print information
		log.Print(bucket.Name)
	}

	log.Println("Successfully listed buckets")
}
