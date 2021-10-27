package marionette

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

var event uint = 0

func GoFishing() {
	log.Println("Go fishing!")

	for {
		if event == 0 {
			cast()
			event = 1
		}

		if event == 1 {
			reel()
			robotgo.Sleep(1)
			if emptyBait() {
				break
			}
			event = 0
		}
	}
}

func cast() {
	x, y, c := Get(Find(500))
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

func reel() {
	x, y, c := Get(Find(501))
	for {
		pc := robotgo.GetPixelColor(x, y)
		fmt.Printf("reel: %v\n", pc)
		if strings.HasPrefix(pc, c) {
			robotgo.Move(x, y)
			robotgo.MouseClick("left", false)
			log.Println("reel..")
			break
		}
		robotgo.MilliSleep(250)
	}
}

func emptyBait() bool {
	x, y, c := Get(Find(502))
	return robotgo.GetPixelColor(x, y) == c
}
