// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of posts
	// (GET /api/v1/admin/posts)
	GetAllPosts(c *gin.Context, params GetAllPostsParams)
	// Create a new post
	// (POST /api/v1/admin/posts)
	CreatePost(c *gin.Context)
	// Delete a post
	// (DELETE /api/v1/admin/posts/{slug})
	DeletePost(c *gin.Context, slug string)
	// Retrieve a post
	// (GET /api/v1/admin/posts/{slug})
	GetAnyPost(c *gin.Context, slug string)
	// Update a post
	// (PUT /api/v1/admin/posts/{slug})
	UpdatePost(c *gin.Context, slug string)
	// Publish the draft of a post
	// (POST /api/v1/admin/posts/{slug}/publish)
	PublishPost(c *gin.Context, slug string)
	// Unpublish a post overriding its draft
	// (POST /api/v1/admin/posts/{slug}/unpublish)
	UnpublishPost(c *gin.Context, slug string)
	// Retrieve a list of published posts
	// (GET /api/v1/posts)
	GetPublishedPosts(c *gin.Context, params GetPublishedPostsParams)
	// Retrieve a published post
	// (GET /api/v1/posts/{slug})
	GetPublishedPost(c *gin.Context, slug string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAllPosts operation middleware
func (siw *ServerInterfaceWrapper) GetAllPosts(c *gin.Context) {

	var err error

	c.Set(Admin_authorizerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllPostsParams

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter pageSize: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "nextPageToken" -------------

	err = runtime.BindQueryParameter("form", true, false, "nextPageToken", c.Request.URL.Query(), &params.NextPageToken)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter nextPageToken: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAllPosts(c, params)
}

// CreatePost operation middleware
func (siw *ServerInterfaceWrapper) CreatePost(c *gin.Context) {

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreatePost(c)
}

// DeletePost operation middleware
func (siw *ServerInterfaceWrapper) DeletePost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeletePost(c, slug)
}

// GetAnyPost operation middleware
func (siw *ServerInterfaceWrapper) GetAnyPost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAnyPost(c, slug)
}

// UpdatePost operation middleware
func (siw *ServerInterfaceWrapper) UpdatePost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdatePost(c, slug)
}

// PublishPost operation middleware
func (siw *ServerInterfaceWrapper) PublishPost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PublishPost(c, slug)
}

// UnpublishPost operation middleware
func (siw *ServerInterfaceWrapper) UnpublishPost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UnpublishPost(c, slug)
}

// GetPublishedPosts operation middleware
func (siw *ServerInterfaceWrapper) GetPublishedPosts(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPublishedPostsParams

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter pageSize: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "nextPageToken" -------------

	err = runtime.BindQueryParameter("form", true, false, "nextPageToken", c.Request.URL.Query(), &params.NextPageToken)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter nextPageToken: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPublishedPosts(c, params)
}

// GetPublishedPost operation middleware
func (siw *ServerInterfaceWrapper) GetPublishedPost(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameter("simple", false, "slug", c.Param("slug"), &slug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPublishedPost(c, slug)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/admin/posts", wrapper.GetAllPosts)
	router.POST(options.BaseURL+"/api/v1/admin/posts", wrapper.CreatePost)
	router.DELETE(options.BaseURL+"/api/v1/admin/posts/:slug", wrapper.DeletePost)
	router.GET(options.BaseURL+"/api/v1/admin/posts/:slug", wrapper.GetAnyPost)
	router.PUT(options.BaseURL+"/api/v1/admin/posts/:slug", wrapper.UpdatePost)
	router.POST(options.BaseURL+"/api/v1/admin/posts/:slug/publish", wrapper.PublishPost)
	router.POST(options.BaseURL+"/api/v1/admin/posts/:slug/unpublish", wrapper.UnpublishPost)
	router.GET(options.BaseURL+"/api/v1/posts", wrapper.GetPublishedPosts)
	router.GET(options.BaseURL+"/api/v1/posts/:slug", wrapper.GetPublishedPost)
}

type GetAllPostsRequestObject struct {
	Params GetAllPostsParams
}

type GetAllPostsResponseObject interface {
	VisitGetAllPostsResponse(w http.ResponseWriter) error
}

type GetAllPosts200JSONResponse struct {
	// NextPageToken The next page token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`

	// Posts The list of posts
	Posts []AdminPost `json:"posts"`
}

func (response GetAllPosts200JSONResponse) VisitGetAllPostsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAllPosts500Response struct {
}

func (response GetAllPosts500Response) VisitGetAllPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreatePostRequestObject struct {
	Body *multipart.Reader
}

type CreatePostResponseObject interface {
	VisitCreatePostResponse(w http.ResponseWriter) error
}

type CreatePost201Response struct {
}

func (response CreatePost201Response) VisitCreatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type CreatePost400Response struct {
}

func (response CreatePost400Response) VisitCreatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreatePost500Response struct {
}

func (response CreatePost500Response) VisitCreatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type DeletePostRequestObject struct {
	Slug string `json:"slug"`
}

type DeletePostResponseObject interface {
	VisitDeletePostResponse(w http.ResponseWriter) error
}

type DeletePost204Response struct {
}

func (response DeletePost204Response) VisitDeletePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeletePost404Response struct {
}

func (response DeletePost404Response) VisitDeletePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeletePost500Response struct {
}

func (response DeletePost500Response) VisitDeletePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetAnyPostRequestObject struct {
	Slug string `json:"slug"`
}

type GetAnyPostResponseObject interface {
	VisitGetAnyPostResponse(w http.ResponseWriter) error
}

type GetAnyPost200JSONResponse AdminPost

func (response GetAnyPost200JSONResponse) VisitGetAnyPostResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAnyPost404Response struct {
}

func (response GetAnyPost404Response) VisitGetAnyPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetAnyPost500Response struct {
}

func (response GetAnyPost500Response) VisitGetAnyPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type UpdatePostRequestObject struct {
	Slug string `json:"slug"`
	Body *multipart.Reader
}

type UpdatePostResponseObject interface {
	VisitUpdatePostResponse(w http.ResponseWriter) error
}

type UpdatePost200Response struct {
}

func (response UpdatePost200Response) VisitUpdatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type UpdatePost400Response struct {
}

func (response UpdatePost400Response) VisitUpdatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UpdatePost404Response struct {
}

func (response UpdatePost404Response) VisitUpdatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type UpdatePost500Response struct {
}

func (response UpdatePost500Response) VisitUpdatePostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PublishPostRequestObject struct {
	Slug string `json:"slug"`
}

type PublishPostResponseObject interface {
	VisitPublishPostResponse(w http.ResponseWriter) error
}

type PublishPost201Response struct {
}

func (response PublishPost201Response) VisitPublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PublishPost400Response struct {
}

func (response PublishPost400Response) VisitPublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PublishPost500Response struct {
}

func (response PublishPost500Response) VisitPublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type UnpublishPostRequestObject struct {
	Slug string `json:"slug"`
}

type UnpublishPostResponseObject interface {
	VisitUnpublishPostResponse(w http.ResponseWriter) error
}

type UnpublishPost201Response struct {
}

func (response UnpublishPost201Response) VisitUnpublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type UnpublishPost400Response struct {
}

func (response UnpublishPost400Response) VisitUnpublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UnpublishPost500Response struct {
}

func (response UnpublishPost500Response) VisitUnpublishPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetPublishedPostsRequestObject struct {
	Params GetPublishedPostsParams
}

type GetPublishedPostsResponseObject interface {
	VisitGetPublishedPostsResponse(w http.ResponseWriter) error
}

type GetPublishedPosts200JSONResponse struct {
	// NextPageToken The next page token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`

	// Posts The list of posts
	Posts []Post `json:"posts"`
}

func (response GetPublishedPosts200JSONResponse) VisitGetPublishedPostsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPublishedPosts500Response struct {
}

func (response GetPublishedPosts500Response) VisitGetPublishedPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetPublishedPostRequestObject struct {
	Slug string `json:"slug"`
}

type GetPublishedPostResponseObject interface {
	VisitGetPublishedPostResponse(w http.ResponseWriter) error
}

type GetPublishedPost200JSONResponse Post

func (response GetPublishedPost200JSONResponse) VisitGetPublishedPostResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPublishedPost404Response struct {
}

func (response GetPublishedPost404Response) VisitGetPublishedPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetPublishedPost500Response struct {
}

func (response GetPublishedPost500Response) VisitGetPublishedPostResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Retrieve a list of posts
	// (GET /api/v1/admin/posts)
	GetAllPosts(ctx context.Context, request GetAllPostsRequestObject) (GetAllPostsResponseObject, error)
	// Create a new post
	// (POST /api/v1/admin/posts)
	CreatePost(ctx context.Context, request CreatePostRequestObject) (CreatePostResponseObject, error)
	// Delete a post
	// (DELETE /api/v1/admin/posts/{slug})
	DeletePost(ctx context.Context, request DeletePostRequestObject) (DeletePostResponseObject, error)
	// Retrieve a post
	// (GET /api/v1/admin/posts/{slug})
	GetAnyPost(ctx context.Context, request GetAnyPostRequestObject) (GetAnyPostResponseObject, error)
	// Update a post
	// (PUT /api/v1/admin/posts/{slug})
	UpdatePost(ctx context.Context, request UpdatePostRequestObject) (UpdatePostResponseObject, error)
	// Publish the draft of a post
	// (POST /api/v1/admin/posts/{slug}/publish)
	PublishPost(ctx context.Context, request PublishPostRequestObject) (PublishPostResponseObject, error)
	// Unpublish a post overriding its draft
	// (POST /api/v1/admin/posts/{slug}/unpublish)
	UnpublishPost(ctx context.Context, request UnpublishPostRequestObject) (UnpublishPostResponseObject, error)
	// Retrieve a list of published posts
	// (GET /api/v1/posts)
	GetPublishedPosts(ctx context.Context, request GetPublishedPostsRequestObject) (GetPublishedPostsResponseObject, error)
	// Retrieve a published post
	// (GET /api/v1/posts/{slug})
	GetPublishedPost(ctx context.Context, request GetPublishedPostRequestObject) (GetPublishedPostResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetAllPosts operation middleware
func (sh *strictHandler) GetAllPosts(ctx *gin.Context, params GetAllPostsParams) {
	var request GetAllPostsRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAllPosts(ctx, request.(GetAllPostsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAllPosts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAllPostsResponseObject); ok {
		if err := validResponse.VisitGetAllPostsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreatePost operation middleware
func (sh *strictHandler) CreatePost(ctx *gin.Context) {
	var request CreatePostRequestObject

	if reader, err := ctx.Request.MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.Error(err)
		return
	}

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreatePost(ctx, request.(CreatePostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreatePost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreatePostResponseObject); ok {
		if err := validResponse.VisitCreatePostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeletePost operation middleware
func (sh *strictHandler) DeletePost(ctx *gin.Context, slug string) {
	var request DeletePostRequestObject

	request.Slug = slug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeletePost(ctx, request.(DeletePostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeletePost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(DeletePostResponseObject); ok {
		if err := validResponse.VisitDeletePostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetAnyPost operation middleware
func (sh *strictHandler) GetAnyPost(ctx *gin.Context, slug string) {
	var request GetAnyPostRequestObject

	request.Slug = slug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAnyPost(ctx, request.(GetAnyPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAnyPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAnyPostResponseObject); ok {
		if err := validResponse.VisitGetAnyPostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdatePost operation middleware
func (sh *strictHandler) UpdatePost(ctx *gin.Context, slug string) {
	var request UpdatePostRequestObject

	request.Slug = slug

	if reader, err := ctx.Request.MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.Error(err)
		return
	}

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdatePost(ctx, request.(UpdatePostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdatePost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(UpdatePostResponseObject); ok {
		if err := validResponse.VisitUpdatePostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PublishPost operation middleware
func (sh *strictHandler) PublishPost(ctx *gin.Context, slug string) {
	var request PublishPostRequestObject

	request.Slug = slug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PublishPost(ctx, request.(PublishPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PublishPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PublishPostResponseObject); ok {
		if err := validResponse.VisitPublishPostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UnpublishPost operation middleware
func (sh *strictHandler) UnpublishPost(ctx *gin.Context, slug string) {
	var request UnpublishPostRequestObject

	request.Slug = slug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnpublishPost(ctx, request.(UnpublishPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnpublishPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(UnpublishPostResponseObject); ok {
		if err := validResponse.VisitUnpublishPostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPublishedPosts operation middleware
func (sh *strictHandler) GetPublishedPosts(ctx *gin.Context, params GetPublishedPostsParams) {
	var request GetPublishedPostsRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetPublishedPosts(ctx, request.(GetPublishedPostsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPublishedPosts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetPublishedPostsResponseObject); ok {
		if err := validResponse.VisitGetPublishedPostsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPublishedPost operation middleware
func (sh *strictHandler) GetPublishedPost(ctx *gin.Context, slug string) {
	var request GetPublishedPostRequestObject

	request.Slug = slug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetPublishedPost(ctx, request.(GetPublishedPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPublishedPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetPublishedPostResponseObject); ok {
		if err := validResponse.VisitGetPublishedPostResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
