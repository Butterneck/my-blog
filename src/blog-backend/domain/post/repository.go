package post

import (
	"context"
)

type Repository interface {
	GetAnyPost(ctx context.Context, id string) (*Post, error)
	GetPublishedPost(ctx context.Context, id string) (*Post, error)
	GetAllPosts(ctx context.Context) ([]*Post, error)
	GetPublishedPosts(ctx context.Context) ([]*Post, error)
	CreatePost(ctx context.Context, p *Post) error
	UpdatePost(
		ctx context.Context,
		slug string,
		updateFn func(h *Post) (*Post, error),
	) error
}
