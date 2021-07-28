package marionette

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	event = 0
)

func GoFishing() {
	for {
		if emptyBait(MatchTemplate{Threshold: 0.90}) {
			break
		}

		fishing(MatchTemplate{Threshold: 0.90})
	}

}

func fishing(mt MatchTemplate) {
	if event == 0 {
		c, _ := mt.Find("./assets/cast.png")
		if c.X != 0 {
			robotgo.MouseClick("left", false)
			event = 1
			log.Println("casting..")
		}
	} else {
		r, _ := mt.Find("./assets/reel.png")
		if r.X != 0 {
			robotgo.MouseClick("left", false)
			log.Println("catch!")
			event = 0
			time.Sleep(RandDuration(5, 8))
		}

	}
}

func emptyBait(mt MatchTemplate) bool {
	s, _ := mt.Find("./assets/bait.png")
	if s.X != 0 {
		log.Println("bait empty!")
		return true
	}
	return false
}
