package example1

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

// List is a convenient function for Cobra.
func List(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	// Set the parameters
	params := &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	}

	// List tables
	resp, err := client.ListTables(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	if len(resp.TableNames) == 0 {
		log.Print("no table has been found")
	} else {
		log.Printf("%d tables has been found", len(resp.TableNames))
	}

	// Process the tables
	for _, tableName := range resp.TableNames {
		// Print information
		printInfo(client, tableName)
	}

	log.Println("Successfully listed tables")
}

func printInfo(client *dynamodb.Client, tableName string) {
	// Set the parameters
	params := &dynamodb.DescribeTableInput{
		TableName: &tableName,
	}

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
