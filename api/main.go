package main

import (
	"context"
	"log"
	"os"

	"github.com/Tomdango/retrotool-go/db"
	"github.com/Tomdango/retrotool-go/server"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginRouter *gin.Engine
var ginLambda *ginadapter.GinLambdaV2

func init() {
	log.Printf("Lambda Cold Start")

	godotenv.Load()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	connection, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to initialise DB connection, %v", err)
	}

	ginRouter = server.NewRouter(cfg, connection)
	ginLambda = ginadapter.NewV2(ginRouter)

}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("ENVIRONMENT") == "local" {
		ginRouter.Run()
	} else {
		lambda.Start(Handler)
	}
}
