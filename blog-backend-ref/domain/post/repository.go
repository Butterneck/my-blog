package post

import (
	"context"
)

type PaginatedPosts struct {
	Posts         []*Post
	NextPageToken string
}

type Repository interface {
	GetAnyPost(ctx context.Context, id string) (*Post, error)
	GetPublishedPost(ctx context.Context, id string) (*Post, error)
	GetAllPosts(ctx context.Context, pageSize *int, nextPageToken *string) (*PaginatedPosts, error)
	GetPublishedPosts(ctx context.Context, pageSize *int, nextPageToken *string) (*PaginatedPosts, error)
	CreatePost(ctx context.Context, p *Post) error
	UpdatePost(
		ctx context.Context,
		slug string,
		updateFn func(h *Post) (*Post, error),
	) error
	DeletePost(ctx context.Context, slug string) error
}
