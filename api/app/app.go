package app

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type App struct {
	CognitoClient   *cognito.CognitoIdentityProvider
	UserPoolID      string
	AppClientID     string
	AppClientSecret string
}
