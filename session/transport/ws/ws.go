package ws

import (
	httpClient "net/http"

	"github.com/Flo4604/go-steam/v4/session/transport"
)

type WebSocketTransport struct {
	transport.Transport

	WebClient *httpClient.Client
}

func (w *WebSocketTransport) sendRequest() {

}
