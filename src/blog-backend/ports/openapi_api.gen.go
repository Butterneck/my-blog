// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieve a list of posts
	// (GET /api/v1/admin/posts)
	GetApiV1AdminPosts(c *gin.Context)
	// Create a new post
	// (POST /api/v1/admin/posts)
	PostApiV1AdminPosts(c *gin.Context)
	// Delete a post
	// (DELETE /api/v1/admin/posts/{postSlug})
	DeleteApiV1AdminPostsPostSlug(c *gin.Context, postSlug string)
	// Retrieve a post
	// (GET /api/v1/admin/posts/{postSlug})
	GetApiV1AdminPostsPostSlug(c *gin.Context, postSlug string)
	// Update a post
	// (PUT /api/v1/admin/posts/{postSlug})
	PutApiV1AdminPostsPostSlug(c *gin.Context, postSlug string)
	// Publish the draft of a post
	// (POST /api/v1/admin/posts/{postSlug}/publish)
	PostApiV1AdminPostsPostSlugPublish(c *gin.Context, postSlug string)
	// Retrieve a list of published posts
	// (GET /api/v1/posts)
	GetApiV1Posts(c *gin.Context)
	// Retrieve a published post
	// (GET /api/v1/posts/{postSlug})
	GetApiV1PostsPostSlug(c *gin.Context, postSlug string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetApiV1AdminPosts operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1AdminPosts(c *gin.Context) {

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1AdminPosts(c)
}

// PostApiV1AdminPosts operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1AdminPosts(c *gin.Context) {

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiV1AdminPosts(c)
}

// DeleteApiV1AdminPostsPostSlug operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1AdminPostsPostSlug(c *gin.Context) {

	var err error

	// ------------- Path parameter "postSlug" -------------
	var postSlug string

	err = runtime.BindStyledParameter("simple", false, "postSlug", c.Param("postSlug"), &postSlug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter postSlug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteApiV1AdminPostsPostSlug(c, postSlug)
}

// GetApiV1AdminPostsPostSlug operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1AdminPostsPostSlug(c *gin.Context) {

	var err error

	// ------------- Path parameter "postSlug" -------------
	var postSlug string

	err = runtime.BindStyledParameter("simple", false, "postSlug", c.Param("postSlug"), &postSlug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter postSlug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1AdminPostsPostSlug(c, postSlug)
}

// PutApiV1AdminPostsPostSlug operation middleware
func (siw *ServerInterfaceWrapper) PutApiV1AdminPostsPostSlug(c *gin.Context) {

	var err error

	// ------------- Path parameter "postSlug" -------------
	var postSlug string

	err = runtime.BindStyledParameter("simple", false, "postSlug", c.Param("postSlug"), &postSlug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter postSlug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutApiV1AdminPostsPostSlug(c, postSlug)
}

// PostApiV1AdminPostsPostSlugPublish operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1AdminPostsPostSlugPublish(c *gin.Context) {

	var err error

	// ------------- Path parameter "postSlug" -------------
	var postSlug string

	err = runtime.BindStyledParameter("simple", false, "postSlug", c.Param("postSlug"), &postSlug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter postSlug: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Admin_authorizerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostApiV1AdminPostsPostSlugPublish(c, postSlug)
}

// GetApiV1Posts operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Posts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1Posts(c)
}

// GetApiV1PostsPostSlug operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1PostsPostSlug(c *gin.Context) {

	var err error

	// ------------- Path parameter "postSlug" -------------
	var postSlug string

	err = runtime.BindStyledParameter("simple", false, "postSlug", c.Param("postSlug"), &postSlug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter postSlug: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1PostsPostSlug(c, postSlug)
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

	router.GET(options.BaseURL+"/api/v1/admin/posts", wrapper.GetApiV1AdminPosts)
	router.POST(options.BaseURL+"/api/v1/admin/posts", wrapper.PostApiV1AdminPosts)
	router.DELETE(options.BaseURL+"/api/v1/admin/posts/:postSlug", wrapper.DeleteApiV1AdminPostsPostSlug)
	router.GET(options.BaseURL+"/api/v1/admin/posts/:postSlug", wrapper.GetApiV1AdminPostsPostSlug)
	router.PUT(options.BaseURL+"/api/v1/admin/posts/:postSlug", wrapper.PutApiV1AdminPostsPostSlug)
	router.POST(options.BaseURL+"/api/v1/admin/posts/:postSlug/publish", wrapper.PostApiV1AdminPostsPostSlugPublish)
	router.GET(options.BaseURL+"/api/v1/posts", wrapper.GetApiV1Posts)
	router.GET(options.BaseURL+"/api/v1/posts/:postSlug", wrapper.GetApiV1PostsPostSlug)
}

type GetApiV1AdminPostsRequestObject struct {
}

type GetApiV1AdminPostsResponseObject interface {
	VisitGetApiV1AdminPostsResponse(w http.ResponseWriter) error
}

type GetApiV1AdminPosts200JSONResponse []AdminPost

func (response GetApiV1AdminPosts200JSONResponse) VisitGetApiV1AdminPostsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1AdminPosts500Response struct {
}

func (response GetApiV1AdminPosts500Response) VisitGetApiV1AdminPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostApiV1AdminPostsRequestObject struct {
	Body *PostApiV1AdminPostsJSONRequestBody
}

type PostApiV1AdminPostsResponseObject interface {
	VisitPostApiV1AdminPostsResponse(w http.ResponseWriter) error
}

type PostApiV1AdminPosts201Response struct {
}

func (response PostApiV1AdminPosts201Response) VisitPostApiV1AdminPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PostApiV1AdminPosts400Response struct {
}

func (response PostApiV1AdminPosts400Response) VisitPostApiV1AdminPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PostApiV1AdminPosts500Response struct {
}

func (response PostApiV1AdminPosts500Response) VisitPostApiV1AdminPostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type DeleteApiV1AdminPostsPostSlugRequestObject struct {
	PostSlug string `json:"postSlug"`
}

type DeleteApiV1AdminPostsPostSlugResponseObject interface {
	VisitDeleteApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error
}

type DeleteApiV1AdminPostsPostSlug204Response struct {
}

func (response DeleteApiV1AdminPostsPostSlug204Response) VisitDeleteApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteApiV1AdminPostsPostSlug404Response struct {
}

func (response DeleteApiV1AdminPostsPostSlug404Response) VisitDeleteApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteApiV1AdminPostsPostSlug500Response struct {
}

func (response DeleteApiV1AdminPostsPostSlug500Response) VisitDeleteApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetApiV1AdminPostsPostSlugRequestObject struct {
	PostSlug string `json:"postSlug"`
}

type GetApiV1AdminPostsPostSlugResponseObject interface {
	VisitGetApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error
}

type GetApiV1AdminPostsPostSlug200JSONResponse AdminPost

func (response GetApiV1AdminPostsPostSlug200JSONResponse) VisitGetApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1AdminPostsPostSlug404Response struct {
}

func (response GetApiV1AdminPostsPostSlug404Response) VisitGetApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetApiV1AdminPostsPostSlug500Response struct {
}

func (response GetApiV1AdminPostsPostSlug500Response) VisitGetApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PutApiV1AdminPostsPostSlugRequestObject struct {
	PostSlug string `json:"postSlug"`
	Body     *PutApiV1AdminPostsPostSlugJSONRequestBody
}

type PutApiV1AdminPostsPostSlugResponseObject interface {
	VisitPutApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error
}

type PutApiV1AdminPostsPostSlug200Response struct {
}

func (response PutApiV1AdminPostsPostSlug200Response) VisitPutApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutApiV1AdminPostsPostSlug400Response struct {
}

func (response PutApiV1AdminPostsPostSlug400Response) VisitPutApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PutApiV1AdminPostsPostSlug404Response struct {
}

func (response PutApiV1AdminPostsPostSlug404Response) VisitPutApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutApiV1AdminPostsPostSlug500Response struct {
}

func (response PutApiV1AdminPostsPostSlug500Response) VisitPutApiV1AdminPostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostApiV1AdminPostsPostSlugPublishRequestObject struct {
	PostSlug string `json:"postSlug"`
}

type PostApiV1AdminPostsPostSlugPublishResponseObject interface {
	VisitPostApiV1AdminPostsPostSlugPublishResponse(w http.ResponseWriter) error
}

type PostApiV1AdminPostsPostSlugPublish201Response struct {
}

func (response PostApiV1AdminPostsPostSlugPublish201Response) VisitPostApiV1AdminPostsPostSlugPublishResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PostApiV1AdminPostsPostSlugPublish400Response struct {
}

func (response PostApiV1AdminPostsPostSlugPublish400Response) VisitPostApiV1AdminPostsPostSlugPublishResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PostApiV1AdminPostsPostSlugPublish500Response struct {
}

func (response PostApiV1AdminPostsPostSlugPublish500Response) VisitPostApiV1AdminPostsPostSlugPublishResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetApiV1PostsRequestObject struct {
}

type GetApiV1PostsResponseObject interface {
	VisitGetApiV1PostsResponse(w http.ResponseWriter) error
}

type GetApiV1Posts200JSONResponse []Post

func (response GetApiV1Posts200JSONResponse) VisitGetApiV1PostsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1Posts500Response struct {
}

func (response GetApiV1Posts500Response) VisitGetApiV1PostsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetApiV1PostsPostSlugRequestObject struct {
	PostSlug string `json:"postSlug"`
}

type GetApiV1PostsPostSlugResponseObject interface {
	VisitGetApiV1PostsPostSlugResponse(w http.ResponseWriter) error
}

type GetApiV1PostsPostSlug200JSONResponse Post

func (response GetApiV1PostsPostSlug200JSONResponse) VisitGetApiV1PostsPostSlugResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1PostsPostSlug404Response struct {
}

func (response GetApiV1PostsPostSlug404Response) VisitGetApiV1PostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetApiV1PostsPostSlug500Response struct {
}

func (response GetApiV1PostsPostSlug500Response) VisitGetApiV1PostsPostSlugResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Retrieve a list of posts
	// (GET /api/v1/admin/posts)
	GetApiV1AdminPosts(ctx context.Context, request GetApiV1AdminPostsRequestObject) (GetApiV1AdminPostsResponseObject, error)
	// Create a new post
	// (POST /api/v1/admin/posts)
	PostApiV1AdminPosts(ctx context.Context, request PostApiV1AdminPostsRequestObject) (PostApiV1AdminPostsResponseObject, error)
	// Delete a post
	// (DELETE /api/v1/admin/posts/{postSlug})
	DeleteApiV1AdminPostsPostSlug(ctx context.Context, request DeleteApiV1AdminPostsPostSlugRequestObject) (DeleteApiV1AdminPostsPostSlugResponseObject, error)
	// Retrieve a post
	// (GET /api/v1/admin/posts/{postSlug})
	GetApiV1AdminPostsPostSlug(ctx context.Context, request GetApiV1AdminPostsPostSlugRequestObject) (GetApiV1AdminPostsPostSlugResponseObject, error)
	// Update a post
	// (PUT /api/v1/admin/posts/{postSlug})
	PutApiV1AdminPostsPostSlug(ctx context.Context, request PutApiV1AdminPostsPostSlugRequestObject) (PutApiV1AdminPostsPostSlugResponseObject, error)
	// Publish the draft of a post
	// (POST /api/v1/admin/posts/{postSlug}/publish)
	PostApiV1AdminPostsPostSlugPublish(ctx context.Context, request PostApiV1AdminPostsPostSlugPublishRequestObject) (PostApiV1AdminPostsPostSlugPublishResponseObject, error)
	// Retrieve a list of published posts
	// (GET /api/v1/posts)
	GetApiV1Posts(ctx context.Context, request GetApiV1PostsRequestObject) (GetApiV1PostsResponseObject, error)
	// Retrieve a published post
	// (GET /api/v1/posts/{postSlug})
	GetApiV1PostsPostSlug(ctx context.Context, request GetApiV1PostsPostSlugRequestObject) (GetApiV1PostsPostSlugResponseObject, error)
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

// GetApiV1AdminPosts operation middleware
func (sh *strictHandler) GetApiV1AdminPosts(ctx *gin.Context) {
	var request GetApiV1AdminPostsRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1AdminPosts(ctx, request.(GetApiV1AdminPostsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1AdminPosts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetApiV1AdminPostsResponseObject); ok {
		if err := validResponse.VisitGetApiV1AdminPostsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostApiV1AdminPosts operation middleware
func (sh *strictHandler) PostApiV1AdminPosts(ctx *gin.Context) {
	var request PostApiV1AdminPostsRequestObject

	var body PostApiV1AdminPostsJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiV1AdminPosts(ctx, request.(PostApiV1AdminPostsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiV1AdminPosts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostApiV1AdminPostsResponseObject); ok {
		if err := validResponse.VisitPostApiV1AdminPostsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteApiV1AdminPostsPostSlug operation middleware
func (sh *strictHandler) DeleteApiV1AdminPostsPostSlug(ctx *gin.Context, postSlug string) {
	var request DeleteApiV1AdminPostsPostSlugRequestObject

	request.PostSlug = postSlug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteApiV1AdminPostsPostSlug(ctx, request.(DeleteApiV1AdminPostsPostSlugRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteApiV1AdminPostsPostSlug")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(DeleteApiV1AdminPostsPostSlugResponseObject); ok {
		if err := validResponse.VisitDeleteApiV1AdminPostsPostSlugResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetApiV1AdminPostsPostSlug operation middleware
func (sh *strictHandler) GetApiV1AdminPostsPostSlug(ctx *gin.Context, postSlug string) {
	var request GetApiV1AdminPostsPostSlugRequestObject

	request.PostSlug = postSlug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1AdminPostsPostSlug(ctx, request.(GetApiV1AdminPostsPostSlugRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1AdminPostsPostSlug")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetApiV1AdminPostsPostSlugResponseObject); ok {
		if err := validResponse.VisitGetApiV1AdminPostsPostSlugResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutApiV1AdminPostsPostSlug operation middleware
func (sh *strictHandler) PutApiV1AdminPostsPostSlug(ctx *gin.Context, postSlug string) {
	var request PutApiV1AdminPostsPostSlugRequestObject

	request.PostSlug = postSlug

	var body PutApiV1AdminPostsPostSlugJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutApiV1AdminPostsPostSlug(ctx, request.(PutApiV1AdminPostsPostSlugRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutApiV1AdminPostsPostSlug")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PutApiV1AdminPostsPostSlugResponseObject); ok {
		if err := validResponse.VisitPutApiV1AdminPostsPostSlugResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostApiV1AdminPostsPostSlugPublish operation middleware
func (sh *strictHandler) PostApiV1AdminPostsPostSlugPublish(ctx *gin.Context, postSlug string) {
	var request PostApiV1AdminPostsPostSlugPublishRequestObject

	request.PostSlug = postSlug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiV1AdminPostsPostSlugPublish(ctx, request.(PostApiV1AdminPostsPostSlugPublishRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiV1AdminPostsPostSlugPublish")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostApiV1AdminPostsPostSlugPublishResponseObject); ok {
		if err := validResponse.VisitPostApiV1AdminPostsPostSlugPublishResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetApiV1Posts operation middleware
func (sh *strictHandler) GetApiV1Posts(ctx *gin.Context) {
	var request GetApiV1PostsRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1Posts(ctx, request.(GetApiV1PostsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1Posts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetApiV1PostsResponseObject); ok {
		if err := validResponse.VisitGetApiV1PostsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetApiV1PostsPostSlug operation middleware
func (sh *strictHandler) GetApiV1PostsPostSlug(ctx *gin.Context, postSlug string) {
	var request GetApiV1PostsPostSlugRequestObject

	request.PostSlug = postSlug

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1PostsPostSlug(ctx, request.(GetApiV1PostsPostSlugRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1PostsPostSlug")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetApiV1PostsPostSlugResponseObject); ok {
		if err := validResponse.VisitGetApiV1PostsPostSlugResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}