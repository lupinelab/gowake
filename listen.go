package gowake

import (
	"net"
)

func Listen(port int) (*net.UDPAddr, string, error) {
	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: port,
	}

	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, "", err
	}
	defer listener.Close()

	magicPacket := MagicPacket{}
	remote := &net.UDPAddr{}
	for {
		_, remote, err = listener.ReadFromUDP(magicPacket[:])
		if err != nil {
			return remote, "", err
		} else {
			return remote, net.HardwareAddr.String(magicPacket[96:]), err
		}
	}
}
