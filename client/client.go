package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting client...")
	saddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6940")
	if err != nil {
		fmt.Println(err)

	}
	sconn, err := net.DialUDP("udp", nil, saddr)
	if err != nil {
		fmt.Println(err)

	}

	defer sconn.Close()

	buffer := make([]byte, 1024)

	for {
		data, _, err := sconn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if data > 0 {
			fmt.Printf("Key: %s\n", fmt.Sprint(data))
		}
	}
}
