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
	for {
		pc := robotgo.GetPixelColor(x, y)
		if strings.HasPrefix(pc, c) {
			robotgo.Move(x, y)
			robotgo.MouseClick("left", false)
			log.Println("casting..")
			break
		}
	}
}

func reel(x, y int, c string) {
	for {

		pc := robotgo.GetPixelColor(x, y)
		if strings.HasPrefix(pc, c) {
			robotgo.Move(x, y)
			robotgo.MouseClick("left", false)
			log.Println("reel..")
			robotgo.Sleep(5)
			break
		}

	}
}

func emptyBait(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}
