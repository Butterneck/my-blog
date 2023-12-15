package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type UnpublishPost struct {
	Slug string
}

type UnpublishPostHandler struct {
	postRepository post.Repository
}

func NewUnpublishPostHandler(postRepository post.Repository) UnpublishPostHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return UnpublishPostHandler{postRepository: postRepository}
}

func (c UnpublishPostHandler) Handle(ctx context.Context, cmd UnpublishPost) error {
	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		p.Unpublish()
		return p, nil
	})
}
