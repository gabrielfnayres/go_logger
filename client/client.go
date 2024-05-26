package main

import (
	"fmt"
	"net"
)

const maxBuffer = 1024

func main() {
	fmt.Println("Starting client...")
	saddr, err := net.ResolveUDPAddr("udp", "localhost:6940")
	if err != nil {
		fmt.Println(err)

	}
	sconn, err := net.ListenUDP("udp", saddr)
	if err != nil {
		fmt.Println(err)

	}

	defer sconn.Close()

	buffer := make([]byte, maxBuffer)

	for {
		data, _, err := sconn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if data > 0 {
			fmt.Printf("Key: %s\n", string(buffer[:data]))
		}
	}
}
