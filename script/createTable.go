package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func createDynamoDBSession2() (*session.Session, error) {
	// Set the AWS region and endpoint for your local DynamoDB instance.
	// Set up AWS session with local DynamoDB credentials
	config := &aws.Config{
		//Region:      aws.String("us-west-2"), // Change this to your desired region
		//Endpoint:    aws.String("http://localhost:8000"), // Change this to your local DynamoDB endpoint
		//Credentials: credentials.NewStaticCredentials("xv51ob", "y5j6ga", ""), // Replace with your credentials
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

func InitialTable() {
	sess, err := createDynamoDBSession2()
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create a DynamoDB client.
	svc := dynamodb.New(sess)
	// Create table Movies
	tableName := "user"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("username"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("username"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err = svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)
}

func main() {
	InitialTable()
}
