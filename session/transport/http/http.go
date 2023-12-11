package http

import (
	httpClient "net/http"

	"github.com/Flo4604/go-steam/v4/session/transport"
)

type HTTPTransport struct {
	transport.Transport

	WebClient *httpClient.Client
}

func (h HTTPTransport) sendRequest() {
}

func (h HTTPTransport) Close() {
}
