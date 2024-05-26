package server

import (
	"fmt"
	"net"
)

func main() {
	// setting address
	addr, err := net.ResolveUDPAddr("udp", "4444")

	if err != nil {
		fmt.Println(err)
		return
	}

	// setting connections

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
		return
	}
	// close connection at the end
	defer conn.Close()

}
