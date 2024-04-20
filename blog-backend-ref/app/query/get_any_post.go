package query

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend-back/domain/post"
)

// Get a published or draft post

type GetAnyPost struct {
	Slug string
}

type GetAnyPostHandler struct {
	postRepository post.Repository
}

func NewGetAnyPostHandler(postRepository post.Repository) GetAnyPostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return GetAnyPostHandler{postRepository: postRepository}
}

func (q GetAnyPostHandler) Handle(ctx context.Context, query GetAnyPost) (*post.Post, error) {
	return q.postRepository.GetAnyPost(ctx, query.Slug)
}
