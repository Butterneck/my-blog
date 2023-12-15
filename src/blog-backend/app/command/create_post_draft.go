package command

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type CreatePostDraft struct {
	Title string
	Body  string
}

type CreatePostDraftHandler struct {
	postRepository post.Repository
}

func NewCreatePostDraftHandler(postRepository post.Repository) CreatePostDraftHandler {
	if postRepository == nil {
		panic("postRepository is nil")
	}

	return CreatePostDraftHandler{postRepository: postRepository}
}

func (c CreatePostDraftHandler) Handle(ctx context.Context, cmd CreatePostDraft) error {
	p, err := post.NewPost(cmd.Title, cmd.Body)
	if err != nil {
		return err
	}

	return c.postRepository.CreatePost(ctx, p)
}
