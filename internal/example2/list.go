package example2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
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

	// Set the parameters
	params := &s3.ListBucketsInput{
		Limit: aws.Int32(5),
	}

	// List tables
	resp, err := client.ListBuckets(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	if len(resp.BucketNames) == 0 {
		log.Print("no table has been found")
	} else {
		log.Printf("%d tables has been found", len(resp.BucketNames))
	}

	// Process the buckets
	for _, bucketName := range resp.BucketNames {
		// Print information
		printInfo(client, bucketName)
	}

	fmt.Println("Successfully listed tables")
}

func printInfo(client *s3.Client, tableName string) {
	// Set the parameters
	params := &s3.DescribeTableInput{
		TableName: &tableName,
	}

	s3.Lis

	// Get table information
	resp, err := client.DescribeTable(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to get table information: %v", err)
	}

	log.Printf("Information of table [%s]", tableName)
	log.Printf("Count of items: %d", resp.Table.ItemCount)
	log.Printf("Size (bytes): %d", resp.Table.TableSizeBytes)
	log.Printf("Status: %s", string(resp.Table.TableStatus))
}
