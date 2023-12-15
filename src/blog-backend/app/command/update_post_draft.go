package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type UpdatePostDraft struct {
	Slug  string
	Title *string
	Body  *string
}

type UpdatePostDraftHandler struct {
	postRepository post.Repository
}

func NewUpdatePostDraftHandler(postRepository post.Repository) UpdatePostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return UpdatePostDraftHandler{postRepository: postRepository}
}

func (c UpdatePostDraftHandler) Handle(ctx context.Context, cmd UpdatePostDraft) error {
	return c.postRepository.UpdatePost(ctx, cmd.Slug, func(p *post.Post) (*post.Post, error) {
		if cmd.Title != nil {
			p.UpdateTitle(*cmd.Title)
		}

		if cmd.Body != nil {
			p.UpdateBody(*cmd.Body)
		}

		return p, nil
	})
}
