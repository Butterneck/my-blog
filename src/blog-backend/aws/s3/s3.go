package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/butterneck/my-blog/src/blog-backend/aws"
	"go.uber.org/zap"
)

var s3Client *s3.Client

func Init(log *zap.SugaredLogger) {
	log.Debugf("s3 package - Init")

	cfg, err := aws.LoadDefaultConfig()

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client = s3.NewFromConfig(cfg)
}

func GetS3() *s3.Client {
	return s3Client
}
