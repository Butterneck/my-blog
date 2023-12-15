package query

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type GetPublishedPostsHandler struct {
	postRepository post.Repository
}

func NewGetPublishedPostsHandler(postRepository post.Repository) GetPublishedPostsHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return GetPublishedPostsHandler{postRepository: postRepository}
}

func (q GetPublishedPostsHandler) Handle(ctx context.Context) ([]*post.Post, error) {
	return q.postRepository.GetPublishedPosts(ctx)
}
