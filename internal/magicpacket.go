package gowake

import (
	"fmt"
	"net"
)

type MagicPacket [102]byte

func BuildPacket(mac string) (packet MagicPacket, err error) {
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		fmt.Println("invalid MAC address")
		return packet, err
	}

	if len(hwAddr) != 6 {
		fmt.Println("invalid MAC address")
		return packet, err
	}

	copy(packet[:], []byte{255, 255, 255, 255, 255, 255})

	offset := 6
	for i := 0; i < 16; i++ {
		copy(packet[offset:], hwAddr[:])
		offset += 6
	}

	return packet, nil
}
