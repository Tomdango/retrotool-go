package main

import (
	"context"
	"log"

	"github.com/Tomdango/retrotool-go/db"
	"github.com/Tomdango/retrotool-go/server"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	log.Printf("Lambda Cold Start")
	ginLambda = ginadapter.NewV2(server.NewRouter())

	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	db.TeamsRepository.Init(awsConfig)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
