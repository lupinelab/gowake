package pkg

import (
	"fmt"
	"net"
)

func ListenMagicPacket(port int) (*net.UDPAddr, string, error) {
	var magicpacket magicPacket

	network := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: port,
	}

	listener, err := net.ListenUDP("udp", &network)
	if err != nil {
		return nil, "", err
	}
	defer listener.Close()

	fmt.Printf("Listening for magic packets on port %d:\n", port)

	for {
		_, remote, err := listener.ReadFromUDP(magicpacket[:])
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		return remote, net.HardwareAddr.String(magicpacket[96:]), err
	}
}
