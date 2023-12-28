package ddb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/butterneck/my-blog/blog-backend/aws"
	"go.uber.org/zap"
)

var db *dynamodb.Client

func Init(log *zap.SugaredLogger) {
	log.Debugf("db package - Init")

	cfg, err := aws.LoadDefaultConfig()

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	db = dynamodb.NewFromConfig(cfg)
}

func GetDB() *dynamodb.Client {
	return db
}
