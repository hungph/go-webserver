package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
	"os"
)

var db *dynamodb.DynamoDB

func CreateDynamoDBClient(region string, endpoint string) (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession()

	if err != nil {
		return nil, err
	}

	dynamoDb := dynamodb.New(sess, aws.NewConfig().WithRegion(region).WithEndpoint(endpoint))

	return dynamoDb, nil
}

func GetDynamoDBClient() *dynamodb.DynamoDB {
	return db
}

func GetTableName() string {
	return os.Getenv("dynamodb_table")
}

func init() {
	godotenv.Load()

	dynamodbEndpoint := os.Getenv("dynamodb_endpoint")
	dynamodbRegion := os.Getenv("dynamodb_region")

	dbUri := fmt.Sprintf("endpoint=%s region=%s table=%s", dynamodbEndpoint, dynamodbRegion, os.Getenv("dynamodb_table"))
	fmt.Println(dbUri)

	dynamoDBClient, err := CreateDynamoDBClient(dynamodbRegion, dynamodbEndpoint)

	if err != nil {
		fmt.Print(err)
	}

	db = dynamoDBClient
}
