package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	/*
		extracts api key from the headers of an HTTP request
		Example
		Authorization: ApiKey {inser apiKey here }
	*/
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found in the header 😞")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header 😞")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return vals[1], nil
}
