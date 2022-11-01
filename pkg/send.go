package pkg

import (
	"fmt"
	"net"
)

func SendMagicPacket(mp magicPacket, port int) {
	conn, _ := net.Dial("udp", fmt.Sprintf("255.255.255.255:%d", port))
	_, err := conn.Write(mp[:])
	if err != nil {
		fmt.Println("Could not send magic packet")
		return
	}
	defer conn.Close()
}
