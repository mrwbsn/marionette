package marionette

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

func GoGardening() {
	for {
		police(MatchTemplate{Threshold: 0.80})
		gather(MatchTemplate{Threshold: 0.88})
	}
}

func gather(mt MatchTemplate) {
	g, _ := mt.Find("./assets/gather.png")
	if g.X != 0 {
		robotgo.MoveMouse(900, 374)
		time.Sleep(1800 * time.Millisecond)
		for i := 0; i < 15; i++ {
			time.Sleep(70 * time.Millisecond)
			robotgo.MouseClick("left", false)
			log.Println("gathering..")
		}
	}

}

func police(mt MatchTemplate) {
	p, _ := mt.Find("./assets/police.png")
	if p.X != 0 {
		time.Sleep(10 * time.Second)
		fmt.Println("Police!")
	}

}
