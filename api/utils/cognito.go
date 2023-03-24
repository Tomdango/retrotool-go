package utils

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func CalculateSecretHash(username string, clientID string, clientSecret string) string {
	hash := hmac.New(sha256.New, []byte(clientSecret))
	hash.Write([]byte(username + clientID))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

type CognitoParameters struct {
	UserPoolEndpoint string
	UserPoolID       string
	ClientID         string
	ClientSecret     string
}

func GetCognitoParameters(cfg aws.Config) CognitoParameters {
	ssmClient := ssm.NewFromConfig(cfg)

	params, err := ssmClient.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
		Path:           aws.String("/prod/cognito/"),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("Unable to retrieve SSM DB parameters, %v", err)
	}

	paramMap := make(map[string]string, len(params.Parameters))
	for _, parameter := range params.Parameters {
		name := strings.TrimPrefix(*parameter.Name, "/prod/cognito/")
		paramMap[name] = *parameter.Value
	}

	return CognitoParameters{
		UserPoolEndpoint: paramMap["user_pool_endpoint"],
		UserPoolID:       paramMap["user_pool_id"],
		ClientID:         paramMap["user_pool_client_id"],
		ClientSecret:     paramMap["user_pool_client_secret"],
	}
}

func GetCognitoKeySet(userPoolID string) (jwk.Set, error) {
	url := "https://cognito-idp.eu-west-2.amazonaws.com/" + userPoolID + "/.well-known/jwks.json"
	return jwk.Fetch(context.TODO(), url)
}

func ParseJWTFromHeaders(header http.Header, userPoolID string) (jwt.Token, error) {
	keySet, err := GetCognitoKeySet(userPoolID)
	if err != nil {
		return nil, err
	}

	jwtString := strings.TrimPrefix(header.Get("Authorization"), "Bearer ")
	verifiedJWT, err := jwt.Parse([]byte(jwtString), jwt.WithKeySet(keySet))

	if err != nil {
		return nil, err
	}

	return verifiedJWT, nil
}
