package client

import (
	"crypto/sha1"
	"encoding/binary"
	httpClient "net/http"
	"net/url"
	"os"

	"github.com/Flo4604/go-steam/v4/session"
	"github.com/Flo4604/go-steam/v4/session/transport"
)

type Client struct {
	Transport transport.Transport

	PlatformType int

	WebClient *httpClient.Client

	UserAgent string
}

func (c *Client) getRSAKey(accountName string) {
	c.Transport.SendRequest()
}

func (c *Client) encryptPassword(accountName string, password string) {
}

type DeviceDetails struct {
	DeviceFriendlyName string
	PlatformType       int
	OSType             int
	GamingDeviceType   int
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
	case session.EAuthTokenPlatformType["SteamClient"]:
		refererQuery := map[string]string{
			"IN_CLIENT":       "true",
			"WEBSITE_ID":      "Client",
			"LOCAL_HOSTNAME":  getSpoofedHostname(),
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
				PlatformType:       session.EAuthTokenPlatformType["SteamClient"],
				OSType:             20,
				GamingDeviceType:   1,
			},
		}

	case session.EAuthTokenPlatformType["WebBrowser"]:

		return &PlatformData{
			WebsiteID: "Community",
			Headers: Headers{
				UserAgent: "Mozilla/5.0 (Windows; U; Windows NT 10.0; en-US; Valve Steam Client/default/1665786434; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
				Origin:    "https://steamcommunity.com",
				Referer:   "https://steamcommunity.com",
			},
			DeviceDetails: DeviceDetails{
				DeviceFriendlyName: "Steam Client WebHelper",
				PlatformType:       session.EAuthTokenPlatformType["WebBrowser"],
			},
		}

	case session.EAuthTokenPlatformType["MobileApp"]:
		return &PlatformData{
			WebsiteID: "mobile",
			Headers: Headers{
				UserAgent: "okhttp/3.12.12",
				Cookie:    "mobileClient=android; mobileClientVersion=777777 3.0.0",
			},
			DeviceDetails: DeviceDetails{
				DeviceFriendlyName: "Galaxy S22",
				PlatformType:       session.EAuthTokenPlatformType["MobileApp"],
				OSType:             -500,
				GamingDeviceType:   528,
			},
		}
	}

	panic("Invalid platform type")
}

func getSpoofedHostname() string {
	hostname, _ := os.Hostname()
	hash := sha1.New()
	hash.Write([]byte(hostname))
	sha1 := hash.Sum(nil)

	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	output := "DESKTOP-"
	for i := 0; i < 7; i++ {
		output += string(chars[binary.BigEndian.Uint16(sha1[i:i+2])%uint16(len(chars))])
	}

	return output
}
