package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Tomdango/retrotool-api-lambda-v1/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func GetAWSConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	return cfg
}

func main() {
	cfg := GetAWSConfig()

	handler := &handler.Handler{
		Clients: &handler.Clients{
			Cognito:  cognitoidentity.NewFromConfig(cfg),
			DynamoDB: dynamodb.NewFromConfig(cfg),
		},
		Config: &handler.Config{
			UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
			AppClientID:     os.Getenv("COGNITO_USER_POOL_CLIENT_ID"),
			AppClientSecret: os.Getenv("COGNITO_USER_POOL_CLIENT_SECRET"),
		},
	}

	http.HandleFunc("/teams/create", handler.CreateTeam)

	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
