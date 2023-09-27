package repository

import (
	"fiber-poc-api/database/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofiber/fiber/v2/log"
)

type UserRepository struct {
	dynamodb *dynamodb.DynamoDB
}

func NewUserRepository(dynamodb *dynamodb.DynamoDB) UserRepository {
	return UserRepository{dynamodb}
}
func (repo UserRepository) GetUserByUsername(username, xRequestId string) (*entity.User, error) {

	tableName := "user"

	//// Create a DescribeTableInput
	//input2 := &dynamodb.DescribeTableInput{
	//	TableName: aws.String(tableName),
	//}
	//
	//// Describe the table
	//result2, err := repo.dynamodb.DescribeTable(input2)
	//if err != nil {
	//	fmt.Println("Error describing table:", err)
	//}
	//
	//// Extract and print the column (attribute) names
	//fmt.Println("Column names (attribute names) for table:", tableName)
	//for _, attr := range result2.Table.AttributeDefinitions {
	//	fmt.Println(*attr.AttributeName)
	//}

	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("#ColumnName = :columnValue"),
		ExpressionAttributeNames: map[string]*string{
			"#ColumnName": aws.String("username"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":columnValue": {
				S: aws.String(username),
			},
		},
	}

	data, err := repo.dynamodb.Query(input)
	if err != nil {
		log.Errorf("[%s] Error querying DynamoDB: %+v", xRequestId, err)
		return nil, err
	}
	var list []entity.User
	for _, item := range data.Items {
		var data entity.User
		if err := dynamodbattribute.UnmarshalMap(item, &data); err != nil {
			log.Errorf("[%s] Error unmarshaling item: %+v", xRequestId, err)
			continue
		}
		list = append(list, data)
	}
	log.Infof("[%s] result Count: %d", xRequestId, *data.Count)
	var user entity.User
	if len(list) > 0 {
		user = list[0]
		return &user, nil
	} else {
		return nil, nil
	}

}

func (repo UserRepository) CreateUser(user *entity.User) error {

	log.Infof("data before Marshall: %+v", user)
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("error marshall item: %s", err)
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("user"),
	}

	_, err = repo.dynamodb.PutItem(input)
	if err != nil {
		log.Errorf("error PutItem error message: %+v", err.Error())
		return err
	}

	return nil
}

func (repo UserRepository) UpdateUser(user *entity.User) error {
	updateExpression := "SET #isDeleted = :isDeleted, #updatedDate = :updatedDate"
	expressionAttributeNames := map[string]*string{
		"#isDeleted":   aws.String("isDeleted"),
		"#updatedDate": aws.String("updatedDate"),
	}
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":isDeleted": {
			S: aws.String("Y"),
		},
		":updatedDate": {
			S: aws.String(user.UpdatedDate),
		},
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String("top"),
			},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		ReturnValues:              aws.String("UPDATED_NEW"), // Return the updated item
	}
	result, err := repo.dynamodb.UpdateItem(input)
	if err != nil {
		log.Errorf("error PutItem error message: %+v", err.Error())
		return err
	}
	log.Infof("result: %+v", result)

	return nil
}
