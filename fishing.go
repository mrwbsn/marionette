package marionette

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

func GoFishing() {
	log.Println("Go fishing!")
	cx, cy, cc := Get(Find(500))
	rx, ry, rc := Get(Find(501))
	ex, ey, ec := Get(Find(502))
	event := 0

	for {
		fmt.Println(event)
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
		fmt.Printf("cast: %v\n", pc)
		if strings.HasPrefix(pc, c) {
			robotgo.Move(x, y)
			robotgo.MouseClick("left", false)
			log.Println("casting..")
			break
		}
		robotgo.MilliSleep(250)
	}
}

func reel(x, y int, c string) {
	for {
		pc := robotgo.GetPixelColor(x, y)
		fmt.Printf("reel: %v\n", pc)
		if strings.HasPrefix(pc, c) || strings.HasPrefix(pc, "a") {
			robotgo.Move(x, y)
			robotgo.MouseClick("left", false)
			log.Println("reel..")
			break
		}
		robotgo.MilliSleep(250)
	}
}

func emptyBait(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}
