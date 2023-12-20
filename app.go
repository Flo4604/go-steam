package steam

import (
	"fmt"
	"time"

	"github.com/Flo4604/go-steam/v5/protocol"
	"github.com/Flo4604/go-steam/v5/protocol/protobuf"
	"github.com/Flo4604/go-steam/v5/protocol/steamlang"
	"google.golang.org/protobuf/proto"
)

type App struct {
	client *Client

	options *AppOptions

	latestChangeNumber uint32
}

type AppOptions struct {
	EnablePicsCache bool

	ChangelistUpdateInterval int // in seconds

	PicsCacheAll bool
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

func (a *App) getChangeListUpdate() {
	if !a.options.EnablePicsCache || a.options.ChangelistUpdateInterval <= 0 {
		println("PICS cache is disabled or update interval is 0, not updating changelist")
		return
	}

	ticker := time.NewTicker(time.Duration(a.options.ChangelistUpdateInterval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				println("Updating changelist")

				data := new(protobuf.CMsgClientPICSChangesSinceRequest)
				data.SendAppInfoChanges = proto.Bool(true)
				data.SendPackageInfoChanges = proto.Bool(true)
				data.SinceChangeNumber = proto.Uint32(20785330)

				a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientPICSChangesSinceRequest, data))

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func (a *App) handlePicsChangesSinceResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto picsChangesSince response!")
		return
	}

	body := new(protobuf.CMsgClientPICSChangesSinceResponse)
	msg := packet.ReadProtoMsg(body)

	// dump the msg
	fmt.Printf("handlePicsChangesSinceResponse %+v\n", msg)
}

func (a *App) handlePicsProductInfoResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto picsProductInfo response!")
		return
	}
}

func (a *App) handlePicsAccessTokenResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto picsAccessToken response!")
		return
	}
}

func (a *App) GetAppInfo(appId uint32) {

}

func (a *App) GetAppChangesSince(changedId uint32) {

}

func (a *App) GetProductInfo() {

}
