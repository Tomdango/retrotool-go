package controllers

import (
	"context"
	"fmt"

	"github.com/Tomdango/retrotool-go/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	CognitoClient *cognitoidentityprovider.Client
	CognitoParams utils.CognitoParameters
}

/**
 * Login
 * POST /auth/login
 */
type LoginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBodyTokens struct {
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int32  `json:"expiresIn"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
}

type LoginResponseBody struct {
	Message string                   `json:"message"`
	Tokens  *LoginResponseBodyTokens `json:"tokens"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body LoginRequestBody
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Request Body, %v", err),
		})
		return
	}

	secretHash := utils.CalculateSecretHash(body.Username, c.CognitoParams.ClientID, c.CognitoParams.ClientSecret)

	initAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME":    body.Username,
			"PASSWORD":    body.Password,
			"SECRET_HASH": secretHash,
		},
		ClientId: aws.String(c.CognitoParams.ClientID),
	}

	result, err := c.CognitoClient.InitiateAuth(context.TODO(), initAuthInput)
	if err != nil {
		ctx.JSONP(401, &ErrorResponseBody{
			Message: "Login Failed",
		})
		return
	}

	ctx.JSONP(200, &LoginResponseBody{
		Message: "Logged In",
		Tokens: &LoginResponseBodyTokens{
			AccessToken:  *result.AuthenticationResult.AccessToken,
			ExpiresIn:    result.AuthenticationResult.ExpiresIn,
			IDToken:      *result.AuthenticationResult.IdToken,
			RefreshToken: *result.AuthenticationResult.RefreshToken,
			TokenType:    *result.AuthenticationResult.TokenType,
		},
	})
}

/**
 * Register
 * Post /auth/register
 */
type RegisterRequestBody struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type RegisterResponseBody struct {
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var body RegisterRequestBody
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Request Body, %v", err),
		})
		return
	}

	secretHash := utils.CalculateSecretHash(body.Username, c.CognitoParams.ClientID, c.CognitoParams.ClientSecret)
	user := &cognitoidentityprovider.SignUpInput{
		Username:   aws.String(body.Username),
		Password:   aws.String(body.Password),
		ClientId:   aws.String(c.CognitoParams.ClientID),
		SecretHash: aws.String(secretHash),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(body.PhoneNumber),
			},
		},
	}

	newUser, err := c.CognitoClient.SignUp(context.TODO(), user)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Failed to create user, %v", err),
		})
		return
	}

	ctx.JSONP(200, &RegisterResponseBody{
		Message: "Successfully registered user",
		UserID:  *newUser.UserSub,
	})
}

/**
 * RegisterConfirmation
 * /auth/confirmation
 */
type OTPRequestBody struct {
	OTP      string `json:"otp" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (c *AuthController) RegisterConfirmation(ctx *gin.Context) {
	var body OTPRequestBody
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Request Body, %v", err),
		})
		return
	}

	secretHash := utils.CalculateSecretHash(body.Username, c.CognitoParams.ClientID, c.CognitoParams.ClientSecret)

	confirmInput := &cognitoidentityprovider.ConfirmSignUpInput{
		ConfirmationCode: aws.String(body.OTP),
		Username:         aws.String(body.Username),
		ClientId:         aws.String(c.CognitoParams.ClientID),
		SecretHash:       aws.String(secretHash),
	}

	_, err = c.CognitoClient.ConfirmSignUp(context.TODO(), confirmInput)

	if err != nil {
		ctx.JSONP(400, &ErrorResponseBody{
			Message: fmt.Sprintf("Failed to confirm registration, %v", err),
		})
		return
	}

	ctx.JSONP(204, nil)
}
