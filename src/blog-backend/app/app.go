package app

import (
	"github.com/butterneck/my-blog/src/blog-backend/app/command"
	"github.com/butterneck/my-blog/src/blog-backend/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	PublishPostDraft command.PublishPostDraftHandler
	CreatePostDraft  command.CreatePostDraftHandler
	UpdatePostDraft  command.UpdatePostDraftHandler
	UnpublishPost    command.UnpublishPostHandler
	DeletePost       command.DeletePostHandler
}

type Queries struct {
	GetPublishedPosts query.GetPublishedPostsHandler
	GetPublishedPost  query.GetPublishedPostHandler
	GetAllPosts       query.GetAllPostsHandler
	GetAnyPost        query.GetAnyPostHandler
}
