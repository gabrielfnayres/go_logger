package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}
