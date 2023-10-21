package steam

import (
	"fmt"
	"time"

	"github.com/Flo4604/go-steam/go-steam/v3/protocol"
	"github.com/Flo4604/go-steam/go-steam/v3/protocol/protobuf"
	"github.com/Flo4604/go-steam/go-steam/v3/protocol/steamlang"
)

type App struct {
	client *Client

	options *AppOptions
}

type AppOptions struct {
	enablePicsCache bool

	changelistUpdateInterval int // in seconds

	picsCacheAll bool
}

func (a *App) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	// log all events
	case steamlang.EMsg_ClientPICSChangesSinceResponse:
		a.handlePicsChangesSinceResponse(packet)
	case steamlang.EMsg_ClientPICSProductInfoResponse:
		a.handlePicsProductInfoResponse(packet)
	}
}

func (a *App) GetAppInfo(appId uint32) {

}

func (a *App) GetAppChangesSince(changedId uint32) {

}

func (a *App) GetProductInfo() {

}

func (a *App) getChangeListUpdate() {
	if !a.options.enablePicsCache || a.options.changelistUpdateInterval <= 0 {
		return
	}

	ticker := time.NewTicker(time.Duration(a.options.changelistUpdateInterval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func (a *App) handlePicsChangesSinceResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}

	body := new(protobuf.CMsgClientPICSChangesSinceResponse)
	msg := packet.ReadProtoMsg(body)

	// dump the msg
	fmt.Printf("%+v\n", msg)
}

func (a *App) handlePicsProductInfoResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}
}

func (a *App) handlePicsAccessTokenResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}
}
