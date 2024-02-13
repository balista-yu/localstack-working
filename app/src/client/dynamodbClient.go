package client

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamodbClient() (*dynamodb.Client, error) {

	log.Printf("endpoint, %v", os.Getenv("AWS_LOCAL_ENDPOINT"))

	awsEndpoint := os.Getenv("AWS_LOCAL_ENDPOINT")
	awsRegion := os.Getenv("AWS_LOCAL_REGION")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}