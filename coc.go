package marionette

import "github.com/go-vgo/robotgo"

func GoCOC() {
	// find coc mission in quest
	x, y, _ := Get(Find(600))
	robotgo.MoveMouse(x, y)
	robotgo.Click("left", false)
}
