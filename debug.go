package steam

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/Flo4604/go-steam/go-steam/v3/protocol"
	"github.com/davecgh/go-spew/spew"
)

type Debug struct {
	client *Client

	packetId, eventId uint64

	options *DebugOptions
}

type DebugOptions struct {
	Enabled bool
	Base    string

	internalBase string
}

func (d *Debug) HandlePacket(packet *protocol.Packet) {
	if !d.options.Enabled {
		return
	}

	d.packetId++
	name := path.Join(d.options.internalBase, "packets", fmt.Sprintf("%d_%d_%s", time.Now().Unix(), d.packetId, packet.EMsg))

	text := packet.String() + "\n\n" + hex.Dump(packet.Data)
	err := os.WriteFile(name+".txt", []byte(text), 0666)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(name+".bin", packet.Data, 0666)
	if err != nil {
		panic(err)
	}
}

func (d *Debug) HandleEvent(event interface{}) {
	if !d.options.Enabled {
		return
	}

	d.eventId++
	name := fmt.Sprintf("%d_%d_%s.txt", time.Now().Unix(), d.eventId, name(event))
	err := os.WriteFile(path.Join(d.options.internalBase, "events", name), []byte(spew.Sdump(event)), 0666)
	if err != nil {
		panic(err)
	}
}

func (d *Debug) init() {
	d.options.internalBase = path.Join(d.options.Base, fmt.Sprint(time.Now().Unix()))

	// clear old files
	err := os.RemoveAll(d.options.internalBase)

	if err != nil {
		err := fmt.Errorf("error clearing old debug files: %v", err)
		fmt.Println(err.Error())
		return
	}

	err = os.MkdirAll(path.Join(d.options.internalBase, "events"), 0700)

	if err != nil {
		err := fmt.Errorf("error creating events directory: %v", err)
		fmt.Println(err.Error())
		return
	}

	err = os.MkdirAll(path.Join(d.options.internalBase, "packets"), 0700)

	if err != nil {
		err := fmt.Errorf("error creating packets directory: %v", err)
		fmt.Println(err.Error())
		return
	}
}

func name(obj interface{}) string {
	val := reflect.ValueOf(obj)
	ind := reflect.Indirect(val)
	if ind.IsValid() {
		return ind.Type().Name()
	} else {
		return val.Type().Name()
	}
}
