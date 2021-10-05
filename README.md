# Golang and AWS

Using Golang to interact with AWS services.

# Prerequisites

## AWS account

You need to create an AWS account.

## Set the configuration

You can either set the configuration with `Environment variables` or `Shared Credentials`, here we'll be using shared credentials files.
To do so, you need to add a file named `credentials` in your home directory.  
Example for Windows, create the file `C:\Users\xxxxx\.aws\credentials` with :

```
[default]
aws_access_key_id = YOUR_KEY
aws_secret_access_key = YOUR_SECRET
region=your-region
```

## Permissions

Here are the permissions needed for this tutorial :
- To successfully complete the PutObject request, you must have the `s3:PutObject` in your IAM permissions.
- To successfully change the objects acl of your PutObject request, you must have the `s3:PutObjectAcl` in your IAM permissions.

# Example 1 : Play with databases (DynamoDB)

In this example we will use DynamoDB databases.

To list tables : `go run main.go example1 list`  
To put stuff into table : `go run main.go example1 put -t yourTableName`

# Example 2 : Play with buckets (S3)

In this example we will use S3 buckets.

To list buckets : `go run main.go example2 list`  
To put stuff into bucket : `go run main.go example2 put -t yourBucketName`