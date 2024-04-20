package query

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend-back/domain/post"
)

// Get a published post

type GetPublishedPost struct {
	PostSlug string
}

type GetPublishedPostHandler struct {
	postRepository post.Repository
}

func NewGetPublishedPostHandler(postRepository post.Repository) GetPublishedPostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return GetPublishedPostHandler{postRepository: postRepository}
}

func (q GetPublishedPostHandler) Handle(ctx context.Context, query GetPublishedPost) (*post.Post, error) {
	return q.postRepository.GetPublishedPost(ctx, query.PostSlug)
}
