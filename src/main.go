package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	_log "github.com/butterneck/my-blog/src/log"
	"github.com/gin-gonic/gin"

	"github.com/butterneck/my-blog/src/db"
)

var ginLambda *ginadapter.GinLambda

func main() {
	log := _log.GetLogger()
	log.Debugf("main package - main")
	db.Init(log)
	lambda.Start(Handler)
}

func init() {
	_log.Init()
	log := _log.GetLogger()
	log.Debugf("main package - init")

	r := gin.Default()
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/posts", GetPosts)
		v1.GET("/posts/:id", GetPost)
		v1.POST("/posts", CreatePost)
		v1.PUT("/posts/:id", UpdatePost)
		v1.DELETE("/posts/:id", DeletePost)

	}

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := _log.GetLogger()

	log.Infof("main package - Handler - request body: %s", req.Body)
	return ginLambda.ProxyWithContext(ctx, req)
}
