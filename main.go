package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			fmt.Println("Event: ", string(ev.Keychar))
		}
	}
}
