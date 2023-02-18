package pkg

import (
	"fmt"
	"net"
)

type SendOptions struct {
	Mp magicPacket
	IP string
	Port int
}

func SendMagicPacket(options SendOptions) {
	conn, _ := net.Dial("udp", fmt.Sprintf("%s:%d", options.IP, options.Port))
	_, err := conn.Write(options.Mp[:])
	if err != nil {
		fmt.Println("Could not send magic packet")
		return
	}
	defer conn.Close()
}
