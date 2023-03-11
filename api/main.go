package main

import (
	"net/http"
	"os"

	"github.com/Tomdango/retrotool-api-lambda-v1/app"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	aws_config := &aws.Config{Region: aws.String("eu-west-2")}
	session, err := session.NewSession(aws_config)

	if err != nil {
		panic(err)
	}

	instance := app.App{
		CognitoClient:   cognito.New(session),
		UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
		AppClientID:     os.Getenv("COGNITO_USER_POOL_CLIENT_ID"),
		AppClientSecret: os.Getenv("COGNITO_USER_POOL_CLIENT_SECRET"),
	}

	http.HandleFunc("/users/register", instance.UserRegisterHandler)
	http.HandleFunc("/users/otp", instance.UserOTPHandler)
	http.HandleFunc("/users/login", instance.UserLoginHandler)

	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
