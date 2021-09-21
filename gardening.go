package marionette

import (
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

var (
	isClicked      bool = false
	clickedCounter int  = 0
)

func GoGardening() {
	DoGardening()
}

func DoGardening() {
	log.Println("Go gathering!")
	rx, ry, rc := Get(Find(101))
	dx, dy, dc := Get(Find(102))
	nx, ny, nc := Get(Find(103))
	gx, gy, gc := Get(Find(201))

	for {
		if day(dx, dy, dc) || night(nx, ny, nc) {
			log.Println("ahh!")
			break
		}

		if captcha(gx, gy, gc) {
			log.Println("police!!")
			robotgo.Sleep(20)
		}

		if !rain(rx, ry, rc) {
			gather(gx, gy, 100, gc)
		} else {
			gather(gx, gy, 3, gc)
		}
	}

}

func gather(x, y, cc int, c string) {
	xc, yc := robotgo.FindColorCS(0xec9c95, 743, 417, 200, 100)

	if isClicked {
		clickedCounter = clickedCounter + 1
	}

	if clickedCounter > cc {
		isClicked = false
		clickedCounter = 0
	}

	if xc == -1 && yc == -1 && !isClicked {
		robotgo.MilliSleep(1700)
		for i := 0; i < 10; i++ {
			robotgo.MilliSleep(200)
			gatherSpam(x, y, c)
		}
		isClicked = true
	}
}

func gatherSpam(x, y int, c string) {
	pc := robotgo.GetPixelColor(x, y)
	// if strings.HasPrefix(pc, c) || strings.HasPrefix(pc, "9") || strings.HasPrefix(pc, "8") {
	if strings.HasPrefix(pc, c) {
		robotgo.MoveMouse(x, y)
		robotgo.MouseClick("left", false)
	}
}

func day(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func night(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func rain(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func captcha(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == "e7e7eb"
}
