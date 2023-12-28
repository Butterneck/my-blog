package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/butterneck/my-blog/blog-backend/aws/ddb"
	"github.com/butterneck/my-blog/blog-backend/aws/s3"
	_log "github.com/butterneck/my-blog/blog-backend/log"
	"github.com/butterneck/my-blog/blog-backend/ports"
	"github.com/butterneck/my-blog/blog-backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	log := _log.GetLogger()
	log.Debugf("main package - main")
	lambda.Start(Handler)
}

// func LogMiddleware(f strictgin.StrictGinHandlerFunc, operationID string) strictgin.StrictGinHandlerFunc {
// 	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
// 		// Your custom middleware logic before handling the request
// 		println("LogMiddleware: Before handling request")
// 		fmt.Println(ctx.Request.Header.Get("Authorization"))

// 		// Call the next middleware or the actual handler
// 		response, err := f(ctx, request)

// 		// Your custom middleware logic after handling the request
// 		println("LogMiddleware: After handling request")

// 		return response, err
// 	}
// }

func init() {
	_log.Init()
	log := _log.GetLogger()
	log.Debugf("main package - init")
	ddb.Init(log)
	s3.Init(log)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default()) // TODO: Remove

	application := service.NewApplication(context.Background())

	httpServer := ports.NewHttpServer(application)
	strictApiHandler := ports.NewStrictHandler(httpServer, nil)
	// strictApiHandler := ports.NewStrictHandler(httpServer, []ports.StrictMiddlewareFunc{
	// 	LogMiddleware,
	// })

	ports.RegisterHandlers(r, strictApiHandler)

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
