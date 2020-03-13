package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const baseURL = "https://api.polygon.io"

var apiKey string

// UseAPIKey ...
func UseAPIKey(key string) {
	apiKey = key
}

// Request ...
func Request(url string, pointer interface{}) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 401:
		return errors.New("unauthorized, check your API key")
	case 404:
		return errors.New("the specified resource was not found")
	case 409:
		return errors.New("parameter is invalid or incorrect")
	}

	return json.NewDecoder(response.Body).Decode(&pointer)
}

// Endpoint ...
func Endpoint(path string, options interface{}) string {
	values, _ := query.Values(options)
	base := fmt.Sprintf("%s%s?apiKey=%s", baseURL, path, apiKey)

	if encoded := values.Encode(); len(encoded) > 0 {
		base = fmt.Sprintf("%s&%s", base, strings.ToLower(encoded))
	}

	return base
}
