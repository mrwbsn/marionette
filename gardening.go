package marionette

import (
	"log"
	"strings"

	"github.com/go-vgo/robotgo"
)

var lock uint = 0

func GoGardening() {
	DoGardening()
}

func DoGardening() {
	log.Println("Go gathering!")

	for {
		if day() || night() {
			log.Println("ahh!")
			break
		}

		if !inGathering() {
			gather()
		}
	}
}

func captchaAnswered() {
	x, y, c := Get(Find(202))
	for {
		pc := robotgo.GetPixelColor(x, y)
		if pc != c {
			log.Println("escape")
			break
		}
	}
}

func gather() {
	x, y, c := Get(Find(200))
	gc := robotgo.GetPixelColor(x, y)

	xc, _ := robotgo.FindColorCS(0xec9c95, 550, 377, 900, 100)
	if xc == -1 {
		lock += 1
	}

	if lock > lockVar() {
		if c != gc && xc == -1 {
			if !captcha() {
				robotgo.MoveMouse(x, y)
				robotgo.MouseClick("left", false)
			} else {
				log.Println("police!!")
				captchaAnswered()
				robotgo.Sleep(1)
			}
		}
		lock = 0
	}

}

func lockVar() uint {
	x, y, c := Get(Find(100))
	gc := robotgo.GetPixelColor(x, y)
	if c == gc {
		return 0
	} else {
		return 5
	}
}

func day() bool {
	x, y, c := Get(Find(101))
	return robotgo.GetPixelColor(x, y) == c
}

func night() bool {
	x, y, c := Get(Find(102))
	return robotgo.GetPixelColor(x, y) == c
}

func captcha() bool {
	x, y, c := Get(Find(201))
	return robotgo.GetPixelColor(x, y) == c
}

func inGathering() bool {
	x, y, c := Get(Find(203))
	pc := robotgo.GetPixelColor(x, y)
	return strings.HasPrefix(pc, c)
}
