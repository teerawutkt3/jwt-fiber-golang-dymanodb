package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createDynamoDBSession() (*session.Session, error) {
	// Set the AWS region and endpoint for your local DynamoDB instance.
	// Set up AWS session with local DynamoDB credentials
	config := &aws.Config{
		Region:      aws.String("us-east-1"), // Change this to your desired region
		//Endpoint:    aws.String("http://localhost:8000"), // Change this to your local DynamoDB endpoint
		Credentials: credentials.NewStaticCredentials("AKIAVU2AOY7ANOZVC2SG", "asfs7aHAO3RgAf3nNY3frfM8ZBOcGiLDQ0O4bYm9", ""), // Replace with your credentials
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func listTables() {
	sess, err := createDynamoDBSession()
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create a DynamoDB client.
	svc := dynamodb.New(sess)

	// List tables in DynamoDB.
	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		fmt.Println("Error listing tables:", err)
		return
	}

	fmt.Println("Tables:")
	for _, tableName := range result.TableNames {
		fmt.Println(*tableName)
	}
}

func main() {
	listTables()
}
