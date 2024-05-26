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
			if ev.Keychar == 32 {
				fmt.Println("Event: Space")
			} else if ev.Keychar == 13 {
				fmt.Println("Event: Return(enter)")
			} else {
				fmt.Println("Event: ", string(ev.Keychar))
			}
		}
	}
}
