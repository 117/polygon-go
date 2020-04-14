// mod github.com/117/polygon@v0.2.1

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

func endpoint(path string, options interface{}) string {
	values, _ := query.Values(options)
	base := fmt.Sprintf("%s%s?apiKey=%s", "https://api.polygon.io", path, os.Getenv("POLYGON_API_KEY"))

	if encoded := values.Encode(); len(encoded) > 0 {
		base = fmt.Sprintf("%s&%s", base, strings.ToLower(encoded))
	}

	return base
}
