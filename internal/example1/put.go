package example1

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

// Put is a convenient function for Cobra.
func Put(cmd *cobra.Command, args []string) {
	// Get the table from flags
	tableName, err := cmd.Flags().GetString("tableName")
	if err != nil {
		log.Fatal("error with tableName flag")
	}

	// Load configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("unable to load configuration: %v", err)
	}

	// Create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	// Set the parameters
	params := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"_id":   &types.AttributeValueMemberS{Value: "12346"},
			"name":  &types.AttributeValueMemberS{Value: "John Doe"},
			"email": &types.AttributeValueMemberS{Value: "john@doe.io"},
		},
	}

	// Put item into table
	_, err = client.PutItem(context.TODO(), params)
	if err != nil {
		log.Fatalf("error while puting item into table: %v", err)
	}

	fmt.Println("Succesfully put item into table")
}
