package gowake

import (
	"fmt"
	"net"
)

func SendMagicPacket(packet MagicPacket, ip string, port int) error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet[:])
	if err != nil {
		return err
	}

	return err
}
