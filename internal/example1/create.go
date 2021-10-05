package example1

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/spf13/cobra"
)

// Create is a convenient function for Cobra.
func Create(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	// Set the parameters
	keyName := "_id"
	name := "test"
	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: &keyName,
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: &keyName,
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName: &name,
	}

	// Create the table
	_, err = client.CreateTable(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to list tables: %v", err)
	}

	log.Println("Successfully created table")
}
