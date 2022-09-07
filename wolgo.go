package main

import (
	"fmt"
	"net"
	"os"
)

type MagicPacket [102]byte

func buildPacket(mac string) (packet MagicPacket, err error) {
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

func sendPacket(mp MagicPacket) {
	conn, _ := net.Dial("udp", "255.255.255.255:9")
	_, err := conn.Write(mp[:])
	if err != nil {
		fmt.Println("Could not send magic packet")
		return
	}
	defer conn.Close()
}

func main() {
	arguements := os.Args
	if len(arguements) == 1 {
		fmt.Println("Please provide a MAC address")
		return
	}
	mp, _ := buildPacket(os.Args[1])
	sendPacket(mp)
}
