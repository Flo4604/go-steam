package session

import (
	"log"

	"github.com/Flo4604/go-steam/v4/session/transport"
	"github.com/Flo4604/go-steam/v4/session/transport/http"
	"github.com/Flo4604/go-steam/v4/session/transport/ws"

	httpClient "net/http"
	"net/url"
)

// types first
var EAuthTokenPlatformType = map[string]int{
	"Unknown":     0,
	"SteamClient": 1,
	"WebBrowser":  2,
	"MobileApp":   3,
}

type Session struct {
	Proxy        string
	PlatformType int

	credentials *CredentialOptions

	transport transport.Transport

	webClient *httpClient.Client
}

func (s *Session) New(prx string, platformType string, transport transport.Transport) {
	s.webClient = &httpClient.Client{}
	s.PlatformType = EAuthTokenPlatformType[platformType]

	if prx != "" {
		proxyUrl, err := url.Parse(prx)

		if err != nil {
			log.Panicln("Error parsing proxy url: ", err)
		}

		s.webClient.Transport = &httpClient.Transport{Proxy: httpClient.ProxyURL(proxyUrl)}
	}

	if transport == nil && platformType == "SteamClient" {
		s.transport = ws.WebSocketTransport{WebClient: s.webClient}
	} else if transport == nil {
		s.transport = http.HTTPTransport{WebClient: s.webClient}
	} else {
		s.transport = transport
	}
}

type CredentialOptions struct {
	AccountName string
	Password    string
	GuardCode   string
}

func (s *Session) StartWithCredentials(opts *CredentialOptions) {
	s.credentials = opts
}

func (s *Session) encryptPassword() {

}

func (s *Session) getRSAKey() {

}

func (s *Session) makeRequest() {

}
