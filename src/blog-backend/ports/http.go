package ports

import (
	"context"
	"io"
	"strings"

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
	resp, err := p.app.Queries.GetAllPosts.Handle(ctx, request.Params.PageSize, request.Params.NextPageToken)

	posts := domainPostsToHttpAdminPosts(resp.Posts)

	return GetAllPosts200JSONResponse{
		Posts:         posts,
		NextPageToken: &resp.NextPageToken,
	}, err
}

func (p HttpServer) CreatePost(ctx context.Context, request CreatePostRequestObject) (CreatePostResponseObject, error) {
	var createPostDraft command.CreatePostDraft

	for {
		part, err := request.Body.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return CreatePost500Response{}, err
		}

		// Check if the part is a file or a form field
		if part.FileName() == "" {
			// It's a form field
			fieldValue, err := io.ReadAll(part)
			if err != nil {
				return CreatePost500Response{}, err
			}

			fieldName := part.FormName()

			if fieldName == "title" {
				createPostDraft.Title = string(fieldValue)
			} else if fieldName == "body" {
				createPostDraft.Body = string(fieldValue)
			}
		} else {
			// It's a file, you can handle it accordingly
			// For example, you can save the file to disk
			fileName := part.FileName()
			file, err := io.ReadAll(part)
			if err != nil {
				return CreatePost500Response{}, err
			}

			createPostDraft.Assets = append(createPostDraft.Assets, command.PostAsset{
				Name: fileName,
				File: file,
			})
		}
	}

	return CreatePost201Response{}, p.app.Commands.CreatePostDraft.Handle(ctx, createPostDraft)
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
	updatePostDraft := command.UpdatePostDraft{
		Slug: request.Slug,
	}

	for {
		part, err := request.Body.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return UpdatePost500Response{}, err
		}

		// Check if the part is a file or a form field
		if part.FileName() == "" {
			// It's a form field
			fieldValue, err := io.ReadAll(part)
			if err != nil {
				return UpdatePost500Response{}, err
			}

			fieldName := part.FormName()

			if fieldName == "title" {
				title := string(fieldValue)
				updatePostDraft.Title = &title
			} else if fieldName == "body" {
				body := string(fieldValue)
				updatePostDraft.Body = &body
			} else if fieldName == "deletedAssets" {
				updatePostDraft.DeletedAssets = strings.Split(string(fieldValue), ",")
			}
		} else {
			// It's a file, you can handle it accordingly
			// For example, you can save the file to disk
			fileName := part.FileName()
			file, err := io.ReadAll(part)
			if err != nil {
				return UpdatePost500Response{}, err
			}

			updatePostDraft.NewAssets = append(updatePostDraft.NewAssets, command.PostAsset{
				Name: fileName,
				File: file,
			})
		}
	}

	return UpdatePost200Response{}, p.app.Commands.UpdatePostDraft.Handle(ctx, updatePostDraft)
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
		Assets:       post.Assets(),
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
		Body:   draft.Body(),
		Title:  draft.Title(),
		Assets: draft.Assets(),
	}
}
