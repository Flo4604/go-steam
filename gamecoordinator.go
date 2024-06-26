package steam

import (
	"bytes"

	"github.com/Flo4604/go-steam/v5/protocol"
	"github.com/Flo4604/go-steam/v5/protocol/gamecoordinator"
	"github.com/Flo4604/go-steam/v5/protocol/protobuf"
	"github.com/Flo4604/go-steam/v5/protocol/steamlang"
	"google.golang.org/protobuf/proto"
)

type GameCoordinator struct {
	client   *Client
	handlers []GCPacketHandler
}

func newGC(client *Client) *GameCoordinator {
	return &GameCoordinator{
		client:   client,
		handlers: make([]GCPacketHandler, 0),
	}
}

type GCPacketHandler interface {
	HandleGCPacket(*gamecoordinator.GCPacket)
}

func (g *GameCoordinator) RegisterPacketHandler(handler GCPacketHandler) {
	g.handlers = append(g.handlers, handler)
}

func (g *GameCoordinator) HandlePacket(packet *protocol.Packet) {
	if packet.EMsg != steamlang.EMsg_ClientFromGC {
		return
	}

	msg := new(protobuf.CMsgGCClient)
	packet.ReadProtoMsg(msg)

	p, err := gamecoordinator.NewGCPacket(msg)
	if err != nil {
		g.client.Errorf("Error reading GC message: %v", err)
		return
	}

	for _, handler := range g.handlers {
		handler.HandleGCPacket(p)
	}
}

func (g *GameCoordinator) Write(msg gamecoordinator.IGCMsg) {
	buf := new(bytes.Buffer)
	msg.Serialize(buf)

	msgType := msg.GetMsgType()
	if msg.IsProto() {
		msgType = msgType | 0x80000000 // mask with protoMask
	}

	g.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientToGC, &protobuf.CMsgGCClient{
		Msgtype: proto.Uint32(msgType),
		Appid:   proto.Uint32(msg.GetAppId()),
		Payload: buf.Bytes(),
	}))
}

// Sets you in the given games. Specify none to quit all games.
func (g *GameCoordinator) SetGamesPlayed(appIds ...uint64) {
	games := make([]*protobuf.CMsgClientGamesPlayed_GamePlayed, 32)
	for _, appId := range appIds {
		// check if we have enough space
		if len(games) == 32 {
			break
		}

		games = append(games, &protobuf.CMsgClientGamesPlayed_GamePlayed{
			GameId: proto.Uint64(appId),
		})
	}

	g.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientGamesPlayed, &protobuf.CMsgClientGamesPlayed{
		GamesPlayed: games,
	}))
}
