package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tomdango/retrotool-api-lambda-v1/utils"
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserRegisterSuccessResponse struct{}

type UserRegisterErrorResponse struct {
	Message string `json:"message"`
}

func (a *App) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	phoneNumber := r.Form.Get("phone_number")

	secretHash := utils.CalculateSecretHash(username, a.AppClientID, a.AppClientSecret)

	user := &cognito.SignUpInput{
		Username:   aws.String(username),
		Password:   aws.String(password),
		ClientId:   aws.String(a.AppClientID),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(phoneNumber),
			},
		},
	}

	_, err := a.CognitoClient.SignUp(user)
	if err != nil {
		payload := UserRegisterErrorResponse{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload)
		return
	}

	payload := UserRegisterSuccessResponse{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((payload))
}

func (a *App) UserOTPHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	otp := r.Form.Get("otp")
	username := r.Form.Get("username")

	secretHash := utils.CalculateSecretHash(username, a.AppClientID, a.AppClientSecret)

	user := &cognito.ConfirmSignUpInput{
		ConfirmationCode: aws.String(otp),
		Username:         aws.String(username),
		ClientId:         aws.String(a.AppClientID),
		SecretHash:       aws.String(secretHash),
	}

	_, err := a.CognitoClient.ConfirmSignUp(user)

	if err != nil {
		fmt.Println(err)
		payload := UserRegisterErrorResponse{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload)
		return
	}

	payload := UserRegisterSuccessResponse{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((payload))
}

func (a *App) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	refresh := r.Form.Get("refresh")
	refreshToken := r.Form.Get("refresh_token")

	secretHash := utils.CalculateSecretHash(username, a.AppClientID, a.AppClientSecret)
	flow := aws.String("USER_PASSWORD_AUTH")

	params := map[string]*string{
		"USERNAME":    aws.String(username),
		"PASSWORD":    aws.String(password),
		"SECRET_HASH": aws.String(secretHash),
	}

	if refresh != "" {
		flow = aws.String("REFRESH_TOKEN_AUTH")
		params = map[string]*string{
			"REFRESH_TOKEN": aws.String(refreshToken),
		}
	}

	authTry := &cognito.InitiateAuthInput{
		AuthFlow:       flow,
		AuthParameters: params,
		ClientId:       aws.String(a.AppClientID),
	}

	res, err := a.CognitoClient.InitiateAuth(authTry)

	if err != nil {
		fmt.Println(err)
		payload := UserRegisterErrorResponse{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(payload)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.AuthenticationResult)
}
