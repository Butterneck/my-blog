package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type DeletePost struct {
	Slug string
}

type DeletePostHandler struct {
	postRepository post.Repository
}

func NewDeletePostHandler(postRepository post.Repository) DeletePostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return DeletePostHandler{postRepository: postRepository}
}

func (c DeletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {
	return c.postRepository.DeletePost(ctx, cmd.Slug)
}
