package service

import (
	"context"
	"os"

	"github.com/butterneck/my-blog/src/blog-backend/adapters"
	"github.com/butterneck/my-blog/src/blog-backend/app"
	"github.com/butterneck/my-blog/src/blog-backend/app/command"
	"github.com/butterneck/my-blog/src/blog-backend/app/query"
	"github.com/butterneck/my-blog/src/blog-backend/aws/ddb"
	"github.com/butterneck/my-blog/src/blog-backend/aws/s3"
)

// This function should configure the app.Application struct and return it.
func NewApplication(ctx context.Context) app.Application {

	// DDB
	ddbClient := ddb.GetDB()
	if ddbClient == nil {
		panic("ddbClient is nil")
	}

	dbTable := os.Getenv("DYNAMODB_TABLE_NAME")
	postsListIndexName := os.Getenv("DYNAMODB_TABLE_POSTS_LIST_INDEX_NAME")
	postsListIndexName = "list-index"
	slugIndexName := os.Getenv("DYNAMODB_TABLE_SLUG_INDEX_NAME")
	slugIndexName = "slug-index"

	// S3
	s3Client := s3.GetS3()
	if s3Client == nil {
		panic("s3Client is nil")
	}

	assetBucketName := os.Getenv("S3_ASSETS_BUCKET_NAME")
	assetBucketName = "butterneck-me-blog-assets"

	postRepository := adapters.NewDDBPostRepository(ddbClient, adapters.DDBPostRepositoryConfig{TableName: dbTable, PostsListIndexName: postsListIndexName, SlugIndexName: slugIndexName})
	assetStore := adapters.NewS3AssetStore(s3Client, adapters.S3AssetStoreConfig{BucketName: assetBucketName})

	return app.Application{
		Commands: app.Commands{
			PublishPostDraft: command.NewPublishPostDraftHandler(postRepository, assetStore),
			CreatePostDraft:  command.NewCreatePostDraftHandler(postRepository, assetStore),
			UpdatePostDraft:  command.NewUpdatePostDraftHandler(postRepository, assetStore),
			UnpublishPost:    command.NewUnpublishPostHandler(postRepository, assetStore),
			DeletePost:       command.NewDeletePostHandler(postRepository, assetStore),
		},
		Queries: app.Queries{
			GetPublishedPosts: query.NewGetPublishedPostsHandler(postRepository),
			GetPublishedPost:  query.NewGetPublishedPostHandler(postRepository),
			GetAllPosts:       query.NewGetAllPostsHandler(postRepository),
			GetAnyPost:        query.NewGetAnyPostHandler(postRepository),
		},
	}
}
