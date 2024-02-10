package machine

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
)

func GetSpoofedHostname() string {
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

func GetMachineId(accountName string) []byte {
	var buffer bytes.Buffer

	buffer.Write(byteToBuffer(0))
	buffer.Write(stringToBuffer("MessageObject"))

	buffer.Write(byteToBuffer(1))
	buffer.Write(stringToBuffer("BB3"))
	buffer.Write(stringToBuffer(sha1Hash(fmt.Sprintf("SteamUser Hash BB3 %s", accountName))))

	buffer.Write(byteToBuffer(1))
	buffer.Write(stringToBuffer("FF2"))
	buffer.Write(stringToBuffer(sha1Hash(fmt.Sprintf("SteamUser Hash FF2 %s", accountName))))

	buffer.Write(byteToBuffer(1))
	buffer.Write(stringToBuffer("3B3"))
	buffer.Write(stringToBuffer(sha1Hash(fmt.Sprintf("SteamUser Hash 3B3 %s", accountName))))

	buffer.Write(byteToBuffer(8))
	buffer.Write(byteToBuffer(8))

	return buffer.Bytes()
}

func sha1Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func stringToBuffer(input string) []byte {
	return append([]byte(input), 0)
}

func byteToBuffer(input byte) []byte {
	return []byte{input}
}
