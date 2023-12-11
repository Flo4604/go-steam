package client

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	httpClient "net/http"
	"net/url"

	"github.com/Flo4604/go-steam/v4/protocol"
	"github.com/Flo4604/go-steam/v4/protocol/protobuf"
	"github.com/Flo4604/go-steam/v4/protocol/steamlang"
	"github.com/Flo4604/go-steam/v4/session/machine"
	"github.com/Flo4604/go-steam/v4/session/transport"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Transport transport.Transport

	PlatformType protobuf.EAuthTokenPlatformType

	WebClient *httpClient.Client

	UserAgent string
}

func (c *Client) getRSAKey(accountName string) *protobuf.CAuthentication_GetPasswordRSAPublicKey_Response {
	resp := c.sendRequest(&transport.RequestData{
		Interface: "Authentication",
		Method:    "GetPasswordRSAPublicKey",
		Version:   1,
		Data: protocol.NewClientMsgProtobuf(steamlang.EMsg_ServiceMethodCallFromClientNonAuthed, &protobuf.CAuthentication_GetPasswordRSAPublicKey_Request{
			AccountName: &accountName,
		}),
	})

	body := new(protobuf.CAuthentication_GetPasswordRSAPublicKey_Response)
	resp.ReadProtoMsg(body)

	return body
}

type EncryptPasswordResponse struct {
	Password     string
	KeyTimestamp uint64
}

func encryptPassword(password string, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedPassword, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(password))
	if err != nil {
		return nil, err
	}
	return encryptedPassword, nil
}

func (c *Client) EncryptPassword(accountName, password string) *EncryptPasswordResponse {
	keyInfo := c.getRSAKey(accountName)

	modBytes, err := hex.DecodeString(*keyInfo.PublickeyMod)
	if err != nil {
		fmt.Println("Error decoding modulus hex:", err)
		return nil
	}

	expBytes, err := hex.DecodeString(*keyInfo.PublickeyExp)
	if err != nil {
		fmt.Println("Error decoding exponent hex:", err)
		return nil
	}

	// Create big integers from byte arrays
	modInt := new(big.Int).SetBytes(modBytes)
	expInt := new(big.Int).SetBytes(expBytes)

	// Create an rsa.PublicKey instance
	publicKey := rsa.PublicKey{
		N: modInt,
		E: int(expInt.Int64()),
	}

	encryptedPassword, err := encryptPassword(password, &publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Convert encrypted password to base64
	return &EncryptPasswordResponse{
		Password:     base64.StdEncoding.EncodeToString(encryptedPassword),
		KeyTimestamp: *keyInfo.Timestamp,
	}
}

type DeviceDetails struct {
	DeviceFriendlyName string
	PlatformType       int32
	OSType             int
	GamingDeviceType   uint32
}
type Headers struct {
	UserAgent string
	Origin    string
	Referer   string
	Cookie    string
}

type PlatformData struct {
	WebsiteID     string
	Headers       Headers
	DeviceDetails DeviceDetails
}

func (c *Client) getPlatformData() *PlatformData {
	switch c.PlatformType {
	case protobuf.EAuthTokenPlatformType_k_EAuthTokenPlatformType_SteamClient:
		refererQuery := map[string]string{
			"IN_CLIENT":       "true",
			"WEBSITE_ID":      "Client",
			"LOCAL_HOSTNAME":  machine.GetSpoofedHostname(),
			"WEBAPI_BASE_URL": "https://api.steampowered.com/",
			"STORE_BASE_URL":  "https://store.steampowered.com/",
			"USE_POPUPS":      "true",
			"DEV_MODE":        "false",
			"LANGUAGE":        "english",
			"PLATFORM":        "windows",
			"COUNTRY":         "US",
			"LAUNCHER_TYPE":   "0",
			"IN_LOGIN":        "true",
		}

		params := url.Values{}
		for key, value := range refererQuery {
			params.Add(key, value)
		}

		return &PlatformData{
			WebsiteID: "Unknown",
			Headers: Headers{
				UserAgent: "Mozilla/5.0 (Windows; U; Windows NT 10.0; en-US; Valve Steam Client/default/1665786434; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
				Origin:    "https://steamloopback.host",
				Referer:   "https://steamloopback.host/index.html?" + params.Encode(),
			},
			DeviceDetails: DeviceDetails{
				DeviceFriendlyName: refererQuery["LOCAL_HOSTNAME"],
				PlatformType:       int32(c.PlatformType),
				OSType:             20,
				GamingDeviceType:   1,
			},
		}

	case protobuf.EAuthTokenPlatformType_k_EAuthTokenPlatformType_WebBrowser:

		return &PlatformData{
			WebsiteID: "Community",
			Headers: Headers{
				UserAgent: "Mozilla/5.0 (Windows; U; Windows NT 10.0; en-US; Valve Steam Client/default/1665786434; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
				Origin:    "https://steamcommunity.com",
				Referer:   "https://steamcommunity.com",
			},
			DeviceDetails: DeviceDetails{
				DeviceFriendlyName: "Steam Client WebHelper",
				PlatformType:       int32(c.PlatformType),
			},
		}

	case protobuf.EAuthTokenPlatformType_k_EAuthTokenPlatformType_MobileApp:
		return &PlatformData{
			WebsiteID: "mobile",
			Headers: Headers{
				UserAgent: "okhttp/3.12.12",
				Cookie:    "mobileClient=android; mobileClientVersion=777777 3.0.0",
			},
			DeviceDetails: DeviceDetails{
				DeviceFriendlyName: "Galaxy S22",
				PlatformType:       int32(c.PlatformType),
				OSType:             -500,
				GamingDeviceType:   528,
			},
		}
	}

	log.Panicf("Invalid platform type %+v", c.PlatformType)
	return nil
}

func (c *Client) sendRequest(request *transport.RequestData) *protocol.Packet {
	response := c.Transport.SendRequest(request)

	if response.Err != nil {
		log.Fatalln("Request failed", response.Err.Error())
	}

	return response.Data
}

type StartSessionWithCredentialsRequest struct {
	AccountName       string
	EncryptedPassword string
	KeyTimestamp      uint64
	Persistence       int32
	PlatformType      int32
}

func (c *Client) StartSessionWithCredentials(rq *StartSessionWithCredentialsRequest) *protobuf.CAuthentication_BeginAuthSessionViaCredentials_Response {
	platformData := c.getPlatformData()

	resp := c.sendRequest(&transport.RequestData{
		Interface: "Authentication",
		Method:    "BeginAuthSessionViaCredentials",
		Version:   1,
		Data: protocol.NewClientMsgProtobuf(steamlang.EMsg_ServiceMethodCallFromClientNonAuthed, &protobuf.CAuthentication_BeginAuthSessionViaCredentials_Request{
			AccountName:         proto.String(rq.AccountName),
			EncryptedPassword:   proto.String(rq.EncryptedPassword),
			EncryptionTimestamp: proto.Uint64(rq.KeyTimestamp),
			RememberLogin:       proto.Bool(true),
			WebsiteId:           proto.String(platformData.WebsiteID),
			Persistence:         (*protobuf.ESessionPersistence)(proto.Int32(rq.Persistence)),
			DeviceDetails: &protobuf.CAuthentication_DeviceDetails{
				DeviceFriendlyName: proto.String(platformData.DeviceDetails.DeviceFriendlyName),
				PlatformType:       (*protobuf.EAuthTokenPlatformType)(&platformData.DeviceDetails.PlatformType),
				OsType:             proto.Int32(int32(platformData.DeviceDetails.OSType)),
				GamingDeviceType:   proto.Uint32(platformData.DeviceDetails.GamingDeviceType),
			},
		}),
	})

	body := new(protobuf.CAuthentication_BeginAuthSessionViaCredentials_Response)
	resp.ReadProtoMsg(body)

	return body
}

type GetLoginStatusRequest struct {
	ClientID  uint64
	RequestID []byte
}

func (c *Client) GetLoginStatus(rq *GetLoginStatusRequest) *protobuf.CAuthentication_PollAuthSessionStatus_Response {
	resp := c.sendRequest(&transport.RequestData{
		Interface: "Authentication",
		Method:    "PollAuthSessionStatus",
		Version:   1,
		Data: protocol.NewClientMsgProtobuf(steamlang.EMsg_ServiceMethodCallFromClientNonAuthed, &protobuf.CAuthentication_PollAuthSessionStatus_Request{
			ClientId:  proto.Uint64(rq.ClientID),
			RequestId: rq.RequestID,
		}),
	})

	body := new(protobuf.CAuthentication_PollAuthSessionStatus_Response)
	resp.ReadProtoMsg(body)

	log.Printf("GetLoginStatus: %+v", body)

	return body
}
