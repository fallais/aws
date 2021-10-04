# Golang and AWS

Using Golang to interact with AWS services.

# Prerequisites

We need to create an AWS account.

# Example 1 : Play with databases (DynamoDB)

In this example we will use DynamoDB databases.

To list tables : `go run main.go example1 list`  
To put stuff into table : `go run main.go example1 put -t yourTableName`

# Example 2 : Play with buckets (S3)

In this example we will use S3 buckets.

To list tables : `go run main.go example2 list`  
To put stuff into table : `go run main.go example2 put -t yourTableName`