package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	client "app/internal/client"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	objectKeys, err := getObjectKeys()
	if err != nil {
		log.Fatalf("failed to get objectKeys, %v", err)
	}

	fmt.Println(objectKeys)
}

func getObjectKeys() (string, error) {

	s3Client, err := client.S3Client()

	if err != nil {
		log.Fatalf("failed to load client, %v", err)
		return "", err
	}

	resp, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_S3_WORKING_DEMO_BUCKET")),
	})
	if err != nil {
		log.Fatalf("failed to buckets, %v", err)
		return "", err
	}

	var objectKeys string
	for _, obj := range resp.Contents {
		objectKeys += *obj.Key + "\n"
	}

	return objectKeys, nil
}
