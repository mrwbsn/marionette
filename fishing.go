package marionette

import (
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

func GoFishing() {
	cx, cy, cc := Get(Find(501))
	rx, ry, rc := Get(Find(502))
	ex, ey, ec := Get(Find(503))
	event := 0

	for {
		if event == 0 {
			cast(cx, cy, cc)
			event = 1
		}

		if event == 1 {
			reel(rx, ry, rc)
			if emptyBait(ex, ey, ec) {
				break
			}
			event = 0
		}
	}

}

func cast(x, y int, c string) {
	robotgo.Move(x, y)
	if robotgo.GetMouseColor() == c {
		robotgo.MouseClick("left", false)
		log.Println("casting..")
	}
}

func reel(x, y int, c string) {
	mc := robotgo.GetMouseColor()
	if strings.HasPrefix(mc, "b") {
		robotgo.MouseClick("left", false)
		log.Println("reel..")
		robotgo.Sleep(5)
	}
}

func emptyBait(x, y int, c string) bool {
	robotgo.Move(x, y)
	mc := robotgo.GetMouseColor()
	return mc == c

}
