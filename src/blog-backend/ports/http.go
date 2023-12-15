package ports

import (
	"context"

	"github.com/butterneck/my-blog/src/blog-backend/app"
	"github.com/butterneck/my-blog/src/blog-backend/app/command"
	"github.com/butterneck/my-blog/src/blog-backend/app/query"
	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
)

type AuthorizationHeaderCtxKey struct{}

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (p HttpServer) GetApiV1AdminPosts(ctx context.Context, request GetApiV1AdminPostsRequestObject) (GetApiV1AdminPostsResponseObject, error) {
	resp, err := p.app.Queries.GetAllPosts.Handle(ctx)

	posts := domainPostsToHttpAdminPosts(resp)

	return GetApiV1AdminPosts200JSONResponse(posts), err
}

func (p HttpServer) PostApiV1AdminPosts(ctx context.Context, request PostApiV1AdminPostsRequestObject) (PostApiV1AdminPostsResponseObject, error) {
	return PostApiV1AdminPosts201Response{}, p.app.Commands.CreatePostDraft.Handle(ctx, command.CreatePostDraft{
		Title: request.Body.Title,
		Body:  request.Body.Body,
	})
}

func (p HttpServer) DeleteApiV1AdminPostsPostSlug(ctx context.Context, request DeleteApiV1AdminPostsPostSlugRequestObject) (DeleteApiV1AdminPostsPostSlugResponseObject, error) {
	// TODO: Implement
	return nil, nil
}

func (p HttpServer) GetApiV1AdminPostsPostSlug(ctx context.Context, request GetApiV1AdminPostsPostSlugRequestObject) (GetApiV1AdminPostsPostSlugResponseObject, error) {
	resp, err := p.app.Queries.GetAnyPost.Handle(ctx, query.GetAnyPost{
		PostSlug: request.PostSlug,
	})

	if resp == nil {
		return GetApiV1AdminPostsPostSlug404Response{}, err
	}

	post := domainPostToHttpAdminPost(resp)

	return GetApiV1AdminPostsPostSlug200JSONResponse(post), err
}

func (p HttpServer) PutApiV1AdminPostsPostSlug(ctx context.Context, request PutApiV1AdminPostsPostSlugRequestObject) (PutApiV1AdminPostsPostSlugResponseObject, error) {
	return PutApiV1AdminPostsPostSlug200Response{}, p.app.Commands.UpdatePostDraft.Handle(ctx, command.UpdatePostDraft{
		Slug:  request.PostSlug,
		Title: request.Body.Title,
		Body:  request.Body.Body,
	})
}

func (p HttpServer) PostApiV1AdminPostsPostSlugPublish(ctx context.Context, request PostApiV1AdminPostsPostSlugPublishRequestObject) (PostApiV1AdminPostsPostSlugPublishResponseObject, error) {
	return PostApiV1AdminPostsPostSlugPublish201Response{}, p.app.Commands.PublishPostDraft.Handle(ctx, command.PublishPostDraft{
		Slug: request.PostSlug,
	})
}

func (p HttpServer) GetApiV1Posts(ctx context.Context, request GetApiV1PostsRequestObject) (GetApiV1PostsResponseObject, error) {
	resp, err := p.app.Queries.GetPublishedPosts.Handle(ctx)

	posts := domainPostsToHttpPosts(resp)

	return GetApiV1Posts200JSONResponse(posts), err
}

func (p HttpServer) GetApiV1PostsPostSlug(ctx context.Context, request GetApiV1PostsPostSlugRequestObject) (GetApiV1PostsPostSlugResponseObject, error) {
	resp, err := p.app.Queries.GetPublishedPost.Handle(ctx, query.GetPublishedPost{
		PostSlug: request.PostSlug,
	})

	if resp == nil {
		return GetApiV1PostsPostSlug404Response{}, err
	}

	post := domainPostToHttpPost(resp)

	return GetApiV1PostsPostSlug200JSONResponse(post), err
}

func domainPostToHttpPost(post *post.Post) Post {
	return Post{
		Body:         post.Body(),
		CreationDate: post.CreationDate(),
		Slug:         post.Slug(),
		Title:        post.Title(),
	}
}

func domainPostsToHttpPosts(posts []*post.Post) []Post {
	httpPosts := make([]Post, len(posts))
	for i, post := range posts {
		httpPosts[i] = domainPostToHttpPost(post)
	}
	return httpPosts
}

func domainPostToHttpAdminPost(post *post.Post) AdminPost {
	return AdminPost{
		Body:         post.Body(),
		CreationDate: post.CreationDate(),
		Draft:        domainDraftToHttpDraft(post.Draft()),
		Slug:         post.Slug(),
		Title:        post.Title(),
	}
}

func domainPostsToHttpAdminPosts(posts []*post.Post) []AdminPost {
	httpPosts := make([]AdminPost, len(posts))
	for i, post := range posts {
		httpPosts[i] = domainPostToHttpAdminPost(post)
	}
	return httpPosts
}

func domainDraftToHttpDraft(draft *post.Draft) PostDraft {
	return PostDraft{
		Body:  draft.Body(),
		Title: draft.Title(),
	}
}
