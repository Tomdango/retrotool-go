package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

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
