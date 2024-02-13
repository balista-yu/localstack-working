package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	client "app/src/client"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	tableName, err := getTableNames()
	if err != nil {
		log.Fatalf("failed to get tables, %v", err)
	}

	fmt.Println(tableName)
}

func getTableNames() (string, error) {

	dynamodbClient, err := client.DynamodbClient()

	if err != nil {
		log.Fatalf("failed to load client, %v", err)
		return "", err
	}

	resp, err := dynamodbClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
		return "", err
	}
	return resp.TableNames[0], nil
}
