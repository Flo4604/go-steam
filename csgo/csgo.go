/*
Provides access to CSGO Game Coordinator functionality.
*/
package csgo

import (
	"fmt"
	"time"

	"github.com/Flo4604/go-steam/v4"
	"github.com/Flo4604/go-steam/v4/csgo/protocol/protobuf"
	"github.com/Flo4604/go-steam/v4/protocol/gamecoordinator"
)

const AppId = 730

// To use any methods of this, you'll need to SetPlaying(true) and wait for
// the GCReadyEvent.
type CSGO struct {
	client *steam.Client

	hasGCSession bool
	isIngame     bool
}

// Creates a new CSGO instance and registers it as a packet handler
func New(client *steam.Client) *CSGO {
	t := &CSGO{client, false, false}
	client.GC.RegisterPacketHandler(t)

	return t
}

func (cs *CSGO) SetPlaying(playing bool) {
	if playing {
		cs.client.GC.SetGamesPlayed(AppId)

		// send CMsgClientHello to GC each 30 seconds until hasGCSession is true
		go func() {
			for !cs.hasGCSession {
				fmt.Println("Sending CMsgClientHello to GC")
				cs.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(protobuf.EGCBaseClientMsg_k_EMsgGCClientHello), &protobuf.CMsgClientHello{}))

				// wait 30 seconds
				time.Sleep(30 * time.Second)
			}
		}()

	} else {
		cs.client.GC.SetGamesPlayed()
	}
}

func (cs *CSGO) GetPlayerProfile(accountId uint32) {
	println("Requesting player profile")
	requestLevel := uint32(32)

	cs.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(protobuf.ECsgoGCMsg_k_EMsgGCCStrike15_v2_ClientRequestPlayersProfile), &protobuf.CMsgGCCStrike15V2_ClientRequestPlayersProfile{
		AccountId:    &accountId,
		RequestLevel: &requestLevel,
	}))
}

type GCReadyEvent struct{}

func (cs *CSGO) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}

	baseMsg := protobuf.EGCBaseClientMsg(packet.MsgType)
	switch baseMsg {
	case protobuf.EGCBaseClientMsg_k_EMsgGCClientWelcome:
		cs.handleWelcome(packet)
		return
	case protobuf.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus:
		cs.handleWelcome(packet)
		return
	}

	csMessage := protobuf.ECsgoGCMsg(packet.MsgType)
	switch csMessage {
	case *protobuf.ECsgoGCMsg_k_EMsgGCCStrike15_v2_PlayersProfile.Enum():
		cs.handlePlayerProfile(packet)
		return
	}

	fmt.Printf("Unknown GC message: %s\n", baseMsg)
}

func (cs *CSGO) handleWelcome(packet *gamecoordinator.GCPacket) {
	println("Received GC welcome")
	// the packet's body is pretty useless
	cs.hasGCSession = true
	cs.client.Emit(&GCReadyEvent{})
}

func (cs *CSGO) handlePlayerProfile(packet *gamecoordinator.GCPacket) {
	println("Received player profile")
	body := new(protobuf.CMsgGCCStrike15V2_PlayersProfile)
	packet.ReadProtoMsg(body)

}
