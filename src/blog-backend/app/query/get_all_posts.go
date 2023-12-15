package query

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type GetAllPostsHandler struct {
	postRepository post.Repository
}

func NewGetAllPostsHandler(postRepository post.Repository) GetAllPostsHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return GetAllPostsHandler{postRepository: postRepository}
}

func (q GetAllPostsHandler) Handle(ctx context.Context) ([]*post.Post, error) {
	return q.postRepository.GetAllPosts(ctx)
}