package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
)

var db *dynamodb.Client

func Init(log *zap.SugaredLogger) {
	log.Debugf("db package - Init")

	cfg, err := config.LoadDefaultConfig(context.TODO()) // config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKID", "SECRET_KEY", "TOKEN")),
	// config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
	// 	func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 		return aws.Endpoint{URL: "http://ddb:8000"}, nil
	// 	},
	// )),
	// config.WithRegion("eu-west-1"),

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	db = dynamodb.NewFromConfig(cfg)
}

func GetDB(log *zap.SugaredLogger) *dynamodb.Client {
	log.Debugf("db package - GetDB")
	return db
}
