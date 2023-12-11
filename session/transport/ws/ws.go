package ws

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	httpClient "net/http"
	"time"

	"github.com/Flo4604/go-steam/v4/protocol"
	"github.com/Flo4604/go-steam/v4/protocol/protobuf"
	"github.com/Flo4604/go-steam/v4/protocol/steamlang"
	"github.com/Flo4604/go-steam/v4/session/transport"
	"golang.org/x/sync/semaphore"
	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

var CMCache = expirable.NewLRU[string, []server](0, nil, time.Minute*5)
var CMSemaphore = semaphore.NewWeighted(int64(1))

type server struct {
	Endpoint       string  `json:"endpoint"`
	LegacyEndpoint string  `json:"legacy_endpoint"`
	Type           string  `json:"type"`
	Dc             string  `json:"dc"`
	Realm          string  `json:"realm"`
	Load           int     `json:"load"`
	WtdLoad        float64 `json:"wtd_load"`
}
type cmListResponse struct {
	Response struct {
		ServerList []server `json:"serverlist"`
		Success    bool     `json:"success"`
		Message    string   `json:"message"`
	} `json:"response"`
}

type Job struct {
	JobId  uint64
	Result chan *protocol.Packet
}

type WebSocketTransport struct {
	webClient *httpClient.Client

	websocket *websocket.Conn

	writer *bytes.Buffer

	ctx context.Context

	jobs map[uint64]chan *protocol.Packet
}

type WebSocketTransportOptions struct {
	WebClient *httpClient.Client
	Ctx       context.Context
}

func NewWebSocketTransport(opts *WebSocketTransportOptions) *WebSocketTransport {
	return &WebSocketTransport{
		writer:    new(bytes.Buffer),
		ctx:       opts.Ctx,
		webClient: opts.WebClient,
		jobs:      make(map[uint64]chan *protocol.Packet),
	}
}

func (w *WebSocketTransport) SendRequest(request *transport.RequestData) *transport.Response {
	if w.websocket == nil {
		w.connect()
	}

	if request.Data == nil {
		return &transport.Response{
			Err:  fmt.Errorf("no data"),
			Data: nil,
		}
	}

	if request.Data.GetMsgType() == steamlang.EMsg_ServiceMethodCallFromClientNonAuthed {
		jobIdBuffer := make([]byte, 8)
		_, err := rand.Read(jobIdBuffer)
		if err != nil {
			panic(err)
		}

		jobIdBuffer[0] &= 0x7F // make sure it's always a positive value

		jobId := binary.BigEndian.Uint64(jobIdBuffer)
		request.Data.SetSourceJobId(protocol.JobId(jobId))
		request.Data.SetRealm(1)
		request.Data.SetSteamId(0)
		request.Data.SetSessionId(0)
		request.Data.SetTargetJobName(fmt.Sprintf("%s.%s#%d", request.Interface, request.Method, request.Version))
	}

	err := request.Data.Serialize(w.writer)
	if err != nil {
		w.writer.Reset()
		log.Fatalf("Error serializing message %v: %v", request.Data, err)
		return nil
	}

	err = w.websocket.Write(w.ctx, websocket.MessageBinary, w.writer.Bytes())
	w.writer.Reset()
	if err != nil {
		log.Fatalf("Error sending message %v: %v", request.Data, err)
	}

	// We won't need the response so stop
	if request.Data.GetMsgType() != steamlang.EMsg_ServiceMethodCallFromClientNonAuthed {
		return &transport.Response{
			Err:  nil,
			Data: nil,
		}
	}

	// create new job
	job := &Job{
		JobId:  uint64(request.Data.GetSourceJobId()),
		Result: make(chan *protocol.Packet),
	}

	// wait for result channel to be closed
	defer close(job.Result)

	w.jobs[job.JobId] = job.Result

	// timeout after 1 minute
	ctx, cancel := context.WithTimeout(w.ctx, time.Minute)
	defer cancel()

	select {
	case result := <-job.Result:
		return &transport.Response{
			Data: result,
			Err:  nil,
		}
	case <-ctx.Done():
		return &transport.Response{
			Data: nil,
			Err:  fmt.Errorf("timeout"),
		}
	}
}

func (w *WebSocketTransport) connect() {
	// We are connected already, force close
	if w.websocket != nil {
		println("closing existing websocket")
		w.websocket.CloseNow()
	}

	println("connecting to websocket")

	servers := w.getCMList()

	var validServers []server

	// we only want servers that support websocket and are in the steamglobal realm
	for _, server := range servers {
		if server.Type != "websockets" || server.Realm != "steamglobal" {
			continue
		}

		validServers = append(validServers, server)
	}

	// choose a random server
	rand := rand.Intn(len(validServers))

	server := validServers[rand]

	con, _, err := websocket.Dial(w.ctx, fmt.Sprintf("wss://%s/cmsocket/", server.Endpoint), nil)

	if err != nil {
		println("error connecting to websocket", err.Error())
		return
	}

	w.websocket = con

	go w.readMessages()

	w.sendHello()
}

func (w *WebSocketTransport) readMessages() {
	for {
		_, message, err := w.websocket.Read(w.ctx)

		if err != nil {
			println("error reading message", err.Error())
			return
		}

		msg, err := protocol.NewPacket(message)

		if err != nil {
			println("error parsing message", err.Error())
			continue
		}

		w.handlePacket(msg)
	}
}

func (w *WebSocketTransport) handlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_Multi:
		w.handleMulti(packet)
	case steamlang.EMsg_ServiceMethodResponse:
		// get job from map
		job, ok := w.jobs[uint64(packet.TargetJobId)]

		if !ok {
			log.Printf("Job not found %s", packet.TargetJobId.String())
			return
		}

		// send result to channel
		job <- packet

	case steamlang.EMsg_ClientLogOnResponse:
	}
}

func (w *WebSocketTransport) handleMulti(msg *protocol.Packet) {
	body := new(protobuf.CMsgMulti)
	msg.ReadProtoMsg(body)
	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		r, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			println("error creating gzip reader", err.Error())
			return
		}

		payload, err = io.ReadAll(r)
		if err != nil {
			println("error reading gzip reader", err.Error())
			return
		}
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := protocol.NewPacket(packetData)
		if err != nil {
			log.Printf("Error reading packet in Multi msg %v: %v", msg, err)
			continue
		}
		w.handlePacket(p)
	}
}

func (w *WebSocketTransport) sendHello() {
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientHello, &protobuf.CMsgClientHello{
		ProtocolVersion: proto.Uint32(65580),
	})

	msg.SetSessionId(0)
	msg.SetSteamId(0)

	w.SendRequest(&transport.RequestData{
		Data: msg,
	})
}

func (w *WebSocketTransport) Close() {
	w.websocket.CloseNow()
}

func (w *WebSocketTransport) getCMList() []server {
	ok := CMSemaphore.TryAcquire(1)
	if !ok {
		println("semaphore not acquired")
		return nil
	}

	defer CMSemaphore.Release(1)

	cmList, ok := CMCache.Get("cmList")

	if ok {
		println("cmList from cache")
		return cmList
	}

	cmList = w.fetchCMList()

	CMCache.Add("cmList", cmList)
	return cmList
}

func (w *WebSocketTransport) fetchCMList() []server {
	resp, err := w.webClient.Get("https://api.steampowered.com/ISteamDirectory/GetCMListForConnect/v0001/?cellid=0")

	if err != nil {
		println("1", err.Error())
		return nil
	}

	defer resp.Body.Close()

	var cmList cmListResponse

	err = json.NewDecoder(resp.Body).Decode(&cmList)

	if err != nil {
		println("2", err.Error())
		return nil
	}

	return cmList.Response.ServerList
}
