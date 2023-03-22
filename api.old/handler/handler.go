package handler

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Clients struct {
	DynamoDB *dynamodb.Client
	Cognito  *cognitoidentity.Client
}

type Config struct {
	UserPoolID      string
	AppClientID     string
	AppClientSecret string
}

type Handler struct {
	Clients *Clients
	Config  *Config
}
