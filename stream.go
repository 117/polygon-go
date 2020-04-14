package polygon

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

// WebSocketServerURL null
type WebSocketServerURL string

var (
	Stocks WebSocketServerURL = "wss://socket.polygon.io/stocks"
	Crypto                    = "wss://socket.polygon.io/crypto"
	Forex                     = "wss://socket.polygon.io/forex"
)

// Stream null
func Stream(url WebSocketServerURL, channels []string, handler func(WebSocketEvent)) error {
	connection, err := websocket.Dial(string(url), "", "http://localhost/")

	if err != nil {
		return err
	}

	for _, action := range []string{
		fmt.Sprintf(`{"action":"auth","params":"%s"}`, os.Getenv("POLYGON_API_KEY")),
		fmt.Sprintf(`{"action":"subscribe","params":"%s"}`, strings.Join(channels, ",")),
	} {
		if _, err := connection.Write([]byte(action)); err != nil {
			return err
		}
	}

	for {
		var (
			events []WebSocketEvent
			msg    = make([]byte, 10240)
			n      int
		)

		if n, err = connection.Read(msg); err != nil {
			return err
		}

		if err := json.Unmarshal(msg[:n], &events); err != nil {
			return err
		}

		for _, event := range events {
			handler(event)
		}
	}
}
