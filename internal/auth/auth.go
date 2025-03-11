package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization : ApiKey {yourkeyhere}
func GetApiKey(headers http.Header) (string, error) {
	keyInHeader := headers.Get("Authorization")
	if keyInHeader == "" {
		return "", errors.New("No api key found")
	}
	tokens := strings.Split(keyInHeader, " ")
	if len(tokens) != 2 || tokens[0] != "ApiKey" {
		return "", errors.New("Malformed header")
	}
	return tokens[1], nil
}
