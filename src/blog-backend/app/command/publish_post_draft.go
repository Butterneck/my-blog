package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type PublishPostDraft struct {
	Slug string
}

type PublishPostDraftHandler struct {
	postRepository post.Repository
}

func NewPublishPostDraftHandler(postRepository post.Repository) PublishPostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return PublishPostDraftHandler{postRepository: postRepository}
}

func (c PublishPostDraftHandler) Handle(ctx context.Context, cmd PublishPostDraft) error {
	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		p.Publish()
		return p, nil
	})
}
