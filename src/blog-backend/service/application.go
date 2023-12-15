package service

import (
	"context"
	"os"

	"github.com/butterneck/my-blog/src/blog-backend/adapters"
	"github.com/butterneck/my-blog/src/blog-backend/app"
	"github.com/butterneck/my-blog/src/blog-backend/app/command"
	"github.com/butterneck/my-blog/src/blog-backend/app/query"
	"github.com/butterneck/my-blog/src/blog-backend/ddb"
)

// This function should configure the app.Application struct and return it.
func NewApplication(ctx context.Context) app.Application {
	ddbClient := ddb.GetDB()
	if ddbClient == nil {
		panic("ddbClient is nil")
	}

	dbTable := os.Getenv("DYNAMODB_TABLE_NAME")
	postsListIndexName := os.Getenv("DYNAMODB_TABLE_POSTS_LIST_INDEX_NAME")
	postsListIndexName = "list-index"
	slugIndexName := os.Getenv("DYNAMODB_TABLE_SLUG_INDEX_NAME")
	slugIndexName = "slug-index"

	postRepository := adapters.NewDDBPostRepository(ddbClient, adapters.DDBPostRepositoryConfig{TableName: dbTable, PostsListIndexName: postsListIndexName, SlugIndexName: slugIndexName})

	return app.Application{
		Commands: app.Commands{
			PublishPostDraft: command.NewPublishPostDraftHandler(postRepository),
			CreatePostDraft:  command.NewCreatePostDraftHandler(postRepository),
			UpdatePostDraft:  command.NewUpdatePostDraftHandler(postRepository),
			UnpublishPost:    command.NewUnpublishPostHandler(postRepository),
			// DeletePostDraft:     command.NewDeletePostDraftHandler(postRepository),
		},
		Queries: app.Queries{
			GetPublishedPosts: query.NewGetPublishedPostsHandler(postRepository),
			GetPublishedPost:  query.NewGetPublishedPostHandler(postRepository),
			GetAllPosts:       query.NewGetAllPostsHandler(postRepository),
			GetAnyPost:        query.NewGetAnyPostHandler(postRepository),
		},
	}
}
