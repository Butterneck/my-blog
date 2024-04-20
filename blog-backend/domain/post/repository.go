package post

import "context"

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

type PostRepositoryAdapter struct {
	Title        Title
	Body         Body
	Slug         Slug
	CreationDate int64
	Draft        DraftRepositoryAdapter
	Assets       []string
}

type DraftRepositoryAdapter struct {
	Title  Title    `json:"title"`
	Body   Body     `json:"body"`
	Slug   Slug     `json:"slug"`
	Assets []string `json:"assets"`
}

// This method is used to load objects from repositories through a PostRepositoryAdapter object
// DO not use this method to create a new Post object, it may not be in a valid state, use NewPost instead
func LoadPostFromRepository(post PostRepositoryAdapter) (*Post, error) {
	return &Post{
		title:        post.Title,
		body:         post.Body,
		slug:         post.Slug,
		creationDate: post.CreationDate,
		assets:       post.Assets,
		draft: Draft{
			title:  post.Draft.Title,
			body:   post.Draft.Body,
			slug:   post.Draft.Slug,
			assets: post.Draft.Assets,
		},
	}, nil
}
