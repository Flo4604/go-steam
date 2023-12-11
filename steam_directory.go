package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Flo4604/go-steam/v4/netutil"
)

// Load initial server list from Steam Directory Web API.
// Call InitializeSteamDirectory() before Connect() to use
// steam directory server list instead of static one.
func InitializeSteamDirectory() error {
	return steamDirectoryCache.Initialize()
}

var steamDirectoryCache *steamDirectory = &steamDirectory{}

type steamDirectory struct {
	sync.RWMutex
	servers       []Server
	isInitialized bool
}

type Server struct {
	Endpoint       string  `json:"endpoint"`
	LegacyEndpoint string  `json:"legacy_endpoint"`
	Type           string  `json:"type"`
	Dc             string  `json:"dc"`
	Realm          string  `json:"realm"`
	Load           int     `json:"load"`
	WtdLoad        float64 `json:"wtd_load"`
}

type ServerListResponse struct {
	Response struct {
		ServerList []Server `json:"serverlist"`
		Success    bool     `json:"success"`
		Message    string   `json:"message"`
	} `json:"response"`
}

// Get server list from steam directory and save it for later
func (sd *steamDirectory) Initialize() error {
	sd.Lock()
	defer sd.Unlock()
	client := new(http.Client)
	resp, err := client.Get("https://api.steampowered.com/ISteamDirectory/GetCMListForConnect/v0001/?cellid=0")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r := ServerListResponse{}

	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}

	if !r.Response.Success {
		return fmt.Errorf("failed to get steam directory, result: %v, message: %v\n", r.Response.Success, r.Response.Message)
	}

	if len(r.Response.ServerList) == 0 {
		return fmt.Errorf("steam returned zero servers for steam directory request\n")
	}

	// filter out servers that are not in realm == 'steamglobal' and not netfilter
	var filteredServers []Server
	for _, server := range r.Response.ServerList {
		if server.Realm == "steamglobal" && server.Type == "netfilter" {
			filteredServers = append(filteredServers, server)
		}
	}

	sd.servers = filteredServers
	sd.isInitialized = true
	return nil
}

func (sd *steamDirectory) GetRandomCM() *netutil.PortAddr {
	sd.RLock()
	defer sd.RUnlock()
	if !sd.isInitialized {
		panic("steam directory is not initialized")
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomServer := sd.servers[rng.Int31n(int32(len(sd.servers)))].Endpoint

	split := strings.Split(randomServer, ":")

	ip, _ := net.DefaultResolver.LookupIPAddr(context.Background(), split[0])

	return netutil.ParsePortAddr(fmt.Sprintf("%s:%s", ip[0].String(), split[1]))
}

func (sd *steamDirectory) IsInitialized() bool {
	sd.RLock()
	defer sd.RUnlock()
	isInitialized := sd.isInitialized
	return isInitialized
}
