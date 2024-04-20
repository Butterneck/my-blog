package command

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend/domain/post"
)

type CreatePostDraft struct {
	Title  string
	Body   string
	Assets []string
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

	// TODO: Replace on the body all the occurrencies of the old filename with the new one (the path on s3)

	// Create post
	p, err := post.NewPost(cmd.Title, cmd.Body, cmd.Assets)
	if err != nil {
		return err
	}

	return c.postRepository.CreatePost(ctx, p)
}
