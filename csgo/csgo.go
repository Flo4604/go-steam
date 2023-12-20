/*
Provides access to CSGO Game Coordinator functionality.
*/
package csgo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Flo4604/go-steam/v5"
	"github.com/Flo4604/go-steam/v5/csgo/protocol/protobuf"
	"github.com/Flo4604/go-steam/v5/protocol/gamecoordinator"
)

const AppId = 730

// To use any methods of this, you'll need to SetPlaying(true) and wait for
// the GCReadyEvent.
type CSGO struct {
	client *steam.Client

	hasGCSession bool
	isIngame     bool

	// array of requests which have a accountId and a handle function
	playerProfileChannels sync.Map
}

type GCReadyEvent struct{}
type PlayerProfileEvent = protobuf.CMsgGCCStrike15V2_MatchmakingGC2ClientHello

// Creates a new CSGO instance and registers it as a packet handler
func New(client *steam.Client) *CSGO {
	t := &CSGO{client, false, false, sync.Map{}}
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

func (cs *CSGO) GetPlayerProfile(accountId uint32) (*PlayerProfileEvent, error) {
	requestLevel := uint32(32)

	// Create a context with a timeout of 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Create a unique channel for this request
	responseChannel := make(chan *PlayerProfileEvent)

	// Load existing channels or create a new slice for the account
	value, _ := cs.playerProfileChannels.LoadOrStore(accountId, []chan *PlayerProfileEvent{})
	channels := value.([]chan *PlayerProfileEvent)

	// Add the new response channel to the slice
	channels = append(channels, responseChannel)
	cs.playerProfileChannels.Store(accountId, channels)

	// Send the request via protobuf
	cs.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(protobuf.ECsgoGCMsg_k_EMsgGCCStrike15_v2_ClientRequestPlayersProfile), &protobuf.CMsgGCCStrike15V2_ClientRequestPlayersProfile{
		AccountId:    &accountId,
		RequestLevel: &requestLevel,
	}))

	// Wait for the response or timeout
	select {
	case response := <-responseChannel:
		return response, nil
	case <-ctx.Done():
		// Remove the response channel from the map on timeout
		cs.removeResponseChannel(accountId, responseChannel)
		return nil, fmt.Errorf("request timed out")
	}
}

func (cs *CSGO) removeResponseChannel(accountId uint32, responseChannel chan *PlayerProfileEvent) {
	value, exists := cs.playerProfileChannels.Load(accountId)
	if exists {
		channels := value.([]chan *PlayerProfileEvent)
		for i, ch := range channels {
			if ch == responseChannel {
				// Remove the channel from the slice
				channels = append(channels[:i], channels[i+1:]...)
				cs.playerProfileChannels.Store(accountId, channels)
				close(responseChannel)
				break
			}
		}
	}
}

func (cs *CSGO) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppId {
		return
	}

	baseMsg := protobuf.EGCBaseClientMsg(packet.MsgType)
	fmt.Printf("Unknown GC message: %s\n", baseMsg)

	switch baseMsg {
	case protobuf.EGCBaseClientMsg_k_EMsgGCClientWelcome:
		cs.handleWelcome(packet)
		return
	}

	csMessage := protobuf.ECsgoGCMsg(packet.MsgType)
	switch csMessage {
	case *protobuf.ECsgoGCMsg_k_EMsgGCCStrike15_v2_PlayersProfile.Enum():
		cs.handlePlayerProfile(packet)
		return
	}

}

func (cs *CSGO) handleWelcome(packet *gamecoordinator.GCPacket) {
	// the packet's body is pretty useless
	cs.hasGCSession = true
	cs.client.Emit(&GCReadyEvent{})
}

func (cs *CSGO) handlePlayerProfile(packet *gamecoordinator.GCPacket) {
	println("Received player profile")
	body := new(protobuf.CMsgGCCStrike15V2_PlayersProfile)
	packet.ReadProtoMsg(body)

	// this is a bit hacky but it works
	for _, profile := range body.GetAccountProfiles() {
		// Retrieve the corresponding handle functions using the accountId
		value, exists := cs.playerProfileChannels.Load(profile.GetAccountId())
		if exists {
			channels := value.([]chan *PlayerProfileEvent)
			for _, responseChannel := range channels {
				responseChannel <- profile
			}
		}

	}
}
