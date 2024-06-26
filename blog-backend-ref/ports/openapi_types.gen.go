// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package ports

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	Admin_authorizerScopes = "admin_authorizer.Scopes"
)

// AdminPost defines model for AdminPost.
type AdminPost struct {
	Assets       []string  `json:"assets"`
	Body         string    `json:"body"`
	CreationDate int64     `json:"creationDate"`
	Draft        PostDraft `json:"draft"`
	Slug         string    `json:"slug"`
	Title        string    `json:"title"`
}

// NewPostRequest defines model for NewPostRequest.
type NewPostRequest struct {
	Assets *[]openapi_types.File `json:"assets,omitempty"`
	Body   string                `json:"body"`
	Title  string                `json:"title"`
}

// Post defines model for Post.
type Post struct {
	Body         string `json:"body"`
	CreationDate int64  `json:"creationDate"`
	Slug         string `json:"slug"`
	Title        string `json:"title"`
}

// PostDraft defines model for PostDraft.
type PostDraft struct {
	Assets []string `json:"assets"`
	Body   string   `json:"body"`
	Title  string   `json:"title"`
}

// UpdatePostRequest defines model for UpdatePostRequest.
type UpdatePostRequest struct {
	Body          *string               `json:"body,omitempty"`
	DeletedAssets *[]string             `json:"deletedAssets,omitempty"`
	NewAssets     *[]openapi_types.File `json:"newAssets,omitempty"`
	Title         *string               `json:"title,omitempty"`
}

// GetAllPostsParams defines parameters for GetAllPosts.
type GetAllPostsParams struct {
	// PageSize The numbers of items to return
	PageSize *int `form:"pageSize,omitempty" json:"pageSize,omitempty"`

	// NextPageToken The page token to use for pagination
	NextPageToken *string `form:"nextPageToken,omitempty" json:"nextPageToken,omitempty"`
}

// GetPublishedPostsParams defines parameters for GetPublishedPosts.
type GetPublishedPostsParams struct {
	// PageSize The numbers of items to return
	PageSize *int `form:"pageSize,omitempty" json:"pageSize,omitempty"`

	// NextPageToken The page token to use for pagination
	NextPageToken *string `form:"nextPageToken,omitempty" json:"nextPageToken,omitempty"`
}

// CreatePostMultipartRequestBody defines body for CreatePost for multipart/form-data ContentType.
type CreatePostMultipartRequestBody = NewPostRequest

// UpdatePostMultipartRequestBody defines body for UpdatePost for multipart/form-data ContentType.
type UpdatePostMultipartRequestBody = UpdatePostRequest
