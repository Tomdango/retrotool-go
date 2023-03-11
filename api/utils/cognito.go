package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func CalculateSecretHash(username string, clientID string, clientSecret string) string {
	hash := hmac.New(sha256.New, []byte(clientSecret))
	hash.Write([]byte(username + clientID))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
