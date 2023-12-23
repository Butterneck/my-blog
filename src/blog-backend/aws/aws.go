package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadDefaultConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO()) // config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKID", "SECRET_KEY", "TOKEN")),
	// config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
	// 	func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 		return aws.Endpoint{URL: "http://ddb:8000"}, nil
	// 	},
	// )),
	// config.WithRegion("eu-west-1"),
}
