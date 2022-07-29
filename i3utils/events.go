package i3utils

import "go.i3wm.org/i3/v4"

func WaitForClose() {
	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WindowEvent)
		if ev.Change == "close" {
			break
		}
	}
}
