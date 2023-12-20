package ports

import (
	"context"
	"fmt"

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

func (p HttpServer) GetAllPosts(ctx context.Context, request GetAllPostsRequestObject) (GetAllPostsResponseObject, error) {
	fmt.Println("GetAllPosts")
	fmt.Println(request)
	fmt.Println(request)
	resp, err := p.app.Queries.GetAllPosts.Handle(ctx, request.Params.PageSize, request.Params.NextPageToken)

	posts := domainPostsToHttpAdminPosts(resp.Posts)

	return GetAllPosts200JSONResponse{
		Posts:         posts,
		NextPageToken: &resp.NextPageToken,
	}, err
}

func (p HttpServer) CreatePost(ctx context.Context, request CreatePostRequestObject) (CreatePostResponseObject, error) {
	return CreatePost201Response{}, p.app.Commands.CreatePostDraft.Handle(ctx, command.CreatePostDraft{
		Title: request.Body.Title,
		Body:  request.Body.Body,
	})
}

func (p HttpServer) DeletePost(ctx context.Context, request DeletePostRequestObject) (DeletePostResponseObject, error) {
	// TODO: Implement
	return nil, nil
}

func (p HttpServer) GetAnyPost(ctx context.Context, request GetAnyPostRequestObject) (GetAnyPostResponseObject, error) {
	resp, err := p.app.Queries.GetAnyPost.Handle(ctx, query.GetAnyPost{
		Slug: request.Slug,
	})

	if resp == nil {
		return GetAnyPost404Response{}, err
	}

	post := domainPostToHttpAdminPost(resp)

	return GetAnyPost200JSONResponse(post), err
}

func (p HttpServer) UpdatePost(ctx context.Context, request UpdatePostRequestObject) (UpdatePostResponseObject, error) {
	return UpdatePost200Response{}, p.app.Commands.UpdatePostDraft.Handle(ctx, command.UpdatePostDraft{
		Slug:  request.Slug,
		Title: request.Body.Title,
		Body:  request.Body.Body,
	})
}

func (p HttpServer) PublishPost(ctx context.Context, request PublishPostRequestObject) (PublishPostResponseObject, error) {
	return PublishPost201Response{}, p.app.Commands.PublishPostDraft.Handle(ctx, command.PublishPostDraft{
		Slug: request.Slug,
	})
}

func (p HttpServer) GetPublishedPosts(ctx context.Context, request GetPublishedPostsRequestObject) (GetPublishedPostsResponseObject, error) {
	resp, err := p.app.Queries.GetPublishedPosts.Handle(ctx, request.Params.PageSize, request.Params.NextPageToken)
	if err != nil {
		return nil, err
	}

	posts := domainPostsToHttpPosts(resp.Posts)

	return GetPublishedPosts200JSONResponse{
		Posts:         posts,
		NextPageToken: &resp.NextPageToken,
	}, err
}

func (p HttpServer) GetPublishedPost(ctx context.Context, request GetPublishedPostRequestObject) (GetPublishedPostResponseObject, error) {
	resp, err := p.app.Queries.GetPublishedPost.Handle(ctx, query.GetPublishedPost{
		PostSlug: request.Slug,
	})

	if resp == nil {
		return GetPublishedPost404Response{}, err
	}

	post := domainPostToHttpPost(resp)

	return GetPublishedPost200JSONResponse(post), err
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
