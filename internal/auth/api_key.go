package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	parts := strings.Fields(headers.Get("Authorization"))
	if len(parts) != 2 || parts[0] != "ApiKey" {
		return "", errors.New("missing or invalid API key header")
	}

	return parts[1], nil
}
