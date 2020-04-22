// mod github.com/117/polygon@v0.3.0

package polygon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
)

// BaseURL null
const BaseURL = "https://api.polygon.io"

// DefaultClient null
var DefaultClient = &Client{&Credentials{
	Key: os.Getenv("POLYGON_KEY"),
}}

func request(url string, pointer interface{}) error {
	response, err := http.DefaultClient.Get(url)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New(response.Status)
	}

	return json.NewDecoder(response.Body).Decode(&pointer)
}

func endpoint(credentials *Credentials, path string, options interface{}) string {
	values, _ := query.Values(options)
	base := fmt.Sprintf("%s%s?apiKey=%s", BaseURL, path, credentials.Key)

	if encoded := values.Encode(); len(encoded) > 0 {
		base = fmt.Sprintf("%s&%s", base, strings.ToLower(encoded))
	}

	return base
}
