package query

import (
	"context"

	"github.com/butterneck/my-blog/blog-backend/domain/post"
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

func (q GetPublishedPostsHandler) Handle(ctx context.Context, pageSize *int, nextPageToken *string) (*post.PaginatedPosts, error) {
	return q.postRepository.GetPublishedPosts(ctx, pageSize, nextPageToken)
}
