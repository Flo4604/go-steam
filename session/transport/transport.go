package transport

import (
	"github.com/Flo4604/go-steam/v4/protocol"
	"google.golang.org/protobuf/proto"
)

type RequestData struct {
	Interface   string
	Method      string
	Version     int
	Data        *protocol.ClientMsgProtobuf
	DecodeProto proto.Message
	AccessToken string
}

type Response struct {
	Data *protocol.Packet
	Err  error
}

type Transport interface {
	SendRequest(request *RequestData) *Response
	Close()
}
