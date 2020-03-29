package polygon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// PolygonBaseURL - The base API url for all Polygon.io requests.
const PolygonBaseURL = "https://api.polygon.io"

func init() {
	http.DefaultClient.Timeout = time.Second * 15
}

func req(url string, pointer interface{}) error {
	response, err := http.DefaultClient.Get(url)

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

func end(path string, options interface{}) string {
	values, _ := query.Values(options)
	base := fmt.Sprintf("%s%s?apiKey=%s", PolygonBaseURL, path, os.Getenv("POLYGON_API_KEY"))

	if encoded := values.Encode(); len(encoded) > 0 {
		base = fmt.Sprintf("%s&%s", base, strings.ToLower(encoded))
	}

	return base
}
