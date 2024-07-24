package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the request headers
// and returns it. If the API key is not present, an error
// the Auth header will look like: " Authenticated: API Key ( Insert API Key Here ) "

func GetAPIKey(headers http.Header) (string, error) {

	val := headers.Get("Authenticated")
	if val == "" {
		return "", errors.New("No authetication Info Found !! ")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid Authetication Info Found !! ")
	}
	if vals[0] != "API Key" {
		return "", errors.New("Invalid Authetication Info Found !! ")
	}
	return vals[1], nil

}
