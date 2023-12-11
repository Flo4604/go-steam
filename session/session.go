package session

import (
	"context"
	"log"
	"time"

	"github.com/Flo4604/go-steam/v4/protocol/protobuf"
	"github.com/Flo4604/go-steam/v4/session/client"
	"github.com/Flo4604/go-steam/v4/session/transport"
	"github.com/Flo4604/go-steam/v4/session/transport/http"
	"github.com/Flo4604/go-steam/v4/session/transport/ws"

	httpClient "net/http"
	"net/url"
)

type Session struct {
	Proxy        string
	PlatformType int

	credentials *CredentialOptions
	transport   transport.Transport
	webClient   *httpClient.Client
	authClient  *client.Client

	startResponse   *protobuf.CAuthentication_BeginAuthSessionViaCredentials_Response
	startQrResponse *protobuf.CAuthentication_BeginAuthSessionViaQR_Response

	hadRemoteInteraction bool
	canceledPolling      bool
	pollingTimer         *time.Timer
	pollingStartedTime   *time.Time
	pollingTimeout       *time.Duration
}

type Options struct {
	Proxy        string
	PlatformType string
	Transport    *transport.Transport
	Ctx          context.Context
}

func New(opts *Options) *Session {
	s := &Session{}
	s.webClient = &httpClient.Client{}
	s.PlatformType = int(protobuf.EAuthTokenPlatformType_value[opts.PlatformType])

	if opts.Proxy != "" {
		proxyUrl, err := url.Parse(opts.Proxy)

		if err != nil {
			log.Panicln("Error parsing proxy url: ", err)
		}

		s.webClient.Transport = &httpClient.Transport{Proxy: httpClient.ProxyURL(proxyUrl)}
	}

	if opts.Transport == nil && opts.PlatformType == "k_EAuthTokenPlatformType_SteamClient" {
		s.transport = ws.NewWebSocketTransport(&ws.WebSocketTransportOptions{
			WebClient: s.webClient,
			Ctx:       opts.Ctx,
		})
	} else if opts.Transport == nil {
		s.transport = &http.HTTPTransport{WebClient: s.webClient}
	} else {
		s.transport = *opts.Transport
	}

	s.authClient = &client.Client{
		Transport:    s.transport,
		PlatformType: protobuf.EAuthTokenPlatformType(protobuf.EAuthTokenPlatformType_value[opts.PlatformType]),
		WebClient:    s.webClient,
		UserAgent:    "Steam Client",
	}

	return s
}

type CredentialOptions struct {
	AccountName string
	Password    string
	GuardCode   string
}

type StartWithCredentialsResponse struct {
	IsActionRequired bool
	ValidActions     []StartWithCredentialsOptions
	QrChallengeUrl   string
}

type StartWithCredentialsOptions struct {
	Type   protobuf.EAuthSessionGuardType
	Detail string
}

func (s *Session) StartWithCredentials(opts *CredentialOptions) *StartWithCredentialsResponse {
	s.credentials = opts

	encrypt := s.authClient.EncryptPassword(s.credentials.AccountName, s.credentials.Password)

	s.startResponse = s.authClient.StartSessionWithCredentials(&client.StartSessionWithCredentialsRequest{
		AccountName:       s.credentials.AccountName,
		EncryptedPassword: encrypt.Password,
		KeyTimestamp:      encrypt.KeyTimestamp,
		Persistence:       1,
		PlatformType:      int32(s.PlatformType),
	})

	return s.processStartResponse()
}

func (s *Session) doPoll() {
	if s.canceledPolling {
		return
	}

	// #TODO: clear timer instead
	if s.pollingTimer != nil {
		s.pollingTimer.Stop()
	}

	if s.pollingStartedTime == nil {
		cTime := time.Now()
		s.pollingStartedTime = &cTime
	}

	//TODO: Calculate polling time and stop if time is higher pollingTimeout

	var clientId uint64
	var requestId []byte

	if s.startResponse != nil {
		clientId = s.startResponse.GetClientId()
		requestId = s.startResponse.RequestId
	} else if s.startQrResponse != nil {
		clientId = s.startQrResponse.GetClientId()
		requestId = s.startQrResponse.RequestId
	} else {
		log.Panicln("No start response found")
	}

	resp := s.authClient.GetLoginStatus(&client.GetLoginStatusRequest{
		ClientID:  clientId,
		RequestID: requestId,
	})

	print(resp)
}

func (s *Session) processStartResponse() *StartWithCredentialsResponse {
	s.canceledPolling = false

	var options []StartWithCredentialsOptions

	for _, confirmation := range s.startResponse.AllowedConfirmations {
		cType := confirmation.GetConfirmationType()
		switch cType {
		case protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_None:
			s.doPoll()
			return &StartWithCredentialsResponse{
				IsActionRequired: false,
			}

		case protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_EmailCode:
		case protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_DeviceCode:
			{
				isEmailAuth := cType == protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_EmailCode

				var success bool
				if isEmailAuth {
					success = s.attemptEmailAuth()
				} else {
					success = s.attemptDeviceAuth()
				}

				if success {
					return &StartWithCredentialsResponse{
						IsActionRequired: false,
					}
				}

				options = append(options, StartWithCredentialsOptions{
					Type:   cType,
					Detail: confirmation.GetAssociatedMessage(),
				})
				break
			}
		case protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_EmailConfirmation:
		case protobuf.EAuthSessionGuardType_k_EAuthSessionGuardType_DeviceConfirmation:
			options = append(options, StartWithCredentialsOptions{
				Type:   cType,
				Detail: confirmation.GetAssociatedMessage(),
			})

			s.doPoll()
			break
		default:
			log.Panicln("Unknown confirmation type: ", cType)
		}
	}

	// Not supposed to happen
	if len(options) == 0 {
		log.Panicln("No confirmation options found")
	}

	response := &StartWithCredentialsResponse{
		IsActionRequired: true,
		ValidActions:     options,
	}

	// #TODO: somehow get both responses into the same variable^^
	if s.startQrResponse != nil {
		response.QrChallengeUrl = s.startQrResponse.GetChallengeUrl()
	}

	return response
}

func (s *Session) attemptEmailAuth() bool {
	if s.credentials.GuardCode == "" {
		return false
	}

	return false
}

func (s *Session) attemptDeviceAuth() bool {
	if s.credentials.GuardCode == "" {
		return false
	}

	return false
}
