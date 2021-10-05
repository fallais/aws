package example1

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/spf13/cobra"
)

// Select is a convenient function for Cobra.
func Select(cmd *cobra.Command, args []string) {
	// Get the table from flags
	tableName, err := cmd.Flags().GetString("tableName")
	if err != nil {
		log.Fatal("error with tableName flag")
	}

	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	// Set the parameters
	filterExpression := "contains(#n, :n)"
	params := &dynamodb.ScanInput{
		TableName:        &tableName,
		FilterExpression: &filterExpression,
		ExpressionAttributeNames: map[string]string{
			"#n": "name",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{
				Value: "John",
			},
		},
	}

	// Scan the table
	resp, err := client.Scan(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	if resp.Count == 0 {
		log.Print("no item has been found")
	} else {
		log.Printf("%d items has been found", resp.Count)
	}

	// Process the items
	for _, item := range resp.Items {
		log.Print(item)
	}

	log.Println("Successfully scanned table")
}
