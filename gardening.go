package marionette

import (
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

func GoGardening() {
	DoGardening()
}

func DoGardening() {
	rx, ry, rc := Get(Find(101))
	dx, dy, dc := Get(Find(102))
	nx, ny, nc := Get(Find(103))
	gx, gy, gc := Get(Find(201))
	log.Println("gathering..")

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
			gatherSpam(gx, gy, gc)
		} else {
			gatherSpam(gx, gy, gc)
		}
	}

}

func gatherSpam(x, y int, c string) {
	pc := robotgo.GetPixelColor(x, y)
	if pc == c || strings.HasPrefix(pc, "9") || strings.HasPrefix(pc, "8") {
		robotgo.MouseClick("left", false)
	}
}

func rain(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func day(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func night(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == c
}

func captcha(x, y int, c string) bool {
	return robotgo.GetPixelColor(x, y) == "e7e7eb"
}
