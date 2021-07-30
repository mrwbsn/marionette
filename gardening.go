package marionette

import (
	"log"

	"github.com/go-vgo/robotgo"
)

func GoGardening() {
	DoGardening()
}

func DoGardening() {
	for {
		// t := time.Now().Format("150405")

		isPolice := police(MatchTemplate{Threshold: 0.80})
		if !isPolice {
			gather(MatchTemplate{Threshold: 0.85})
		} else {
			// TODO: solve captcha
			break
		}

	}

}

func gather(mt MatchTemplate) {
	g, _ := mt.Find("./assets/gather.png")
	if g.X != 0 {
		robotgo.MoveMouse(857, 302)
		robotgo.MilliSleep(1500)
		for i := 0; i < 50; i++ {
			robotgo.MouseClick("left", false)
			robotgo.MilliSleep(10)
			log.Println("gathering..")
		}
	}
}

func police(mt MatchTemplate) bool {
	p, _ := mt.Find("./assets/police.png")
	return p.X != 0
}
