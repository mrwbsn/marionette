package marionette

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

var event int = 0

func GoFishing() {
	for {
		isEmpty := emptyBait(MatchTemplate{Threshold: 0.90})
		if !isEmpty {
			fishing(MatchTemplate{Threshold: 0.90})
		} else {
			break
		}
	}

}

func fishing(mt MatchTemplate) {
	if event == 0 {
		c, _ := mt.Find("./assets/cast.png")
		if c.X != 0 {
			robotgo.Move(c.X, c.Y)
			robotgo.MouseClick("left", false)
			event = 1
			log.Println("casting..")
		}
	} else {
		r, _ := mt.Find("./assets/reel.png")
		if r.X != 0 {
			for i := 0; i < 4; i++ {
				time.Sleep(100 * time.Millisecond)
				robotgo.MouseClick("left", false)
				log.Println("reel!")
			}
			log.Println("catch!")
			event = 0
			time.Sleep(RandDuration(5000, 8000))
		}

	}
}

func emptyBait(mt MatchTemplate) bool {
	s, _ := mt.Find("./assets/bait.png")
	return s.X != 0
}
