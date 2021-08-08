package marionette

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-vgo/robotgo"
)

func GoFishing() {
	cx, cy, cc := checkDisplay(FindDisplay(501))
	rx, ry, rc := checkDisplay(FindDisplay(502))
	event := 0

	for {
		if event == 0 {
			cast(cx, cy, cc)
			event = 1
		}

		if event == 1 {
			reel(rx, ry, rc)
			event = 0
		}
	}

}

func checkDisplay(d *Display) (int, int, string) {
	var x, y int
	var c string
	var err error
	if d.Label != "" {
		x, err = strconv.Atoi(d.X)
		if err != nil {
			log.Fatalln(err)
		}

		y, err = strconv.Atoi(d.Y)
		if err != nil {
			log.Fatalln(err)
		}

		c = d.Color
	}
	return x, y, c
}

func cast(x, y int, c string) {
	robotgo.Move(x, y)
	if robotgo.GetMouseColor() == c {
		robotgo.MouseClick("left", false)
		log.Println("casting..")
	}
}

func reel(x, y int, c string) {
	mc := robotgo.GetMouseColor()
	fmt.Println(mc)
	if mc == "b8ea67" {
		robotgo.MouseClick("left", false)
		log.Println("reel..")
		robotgo.Sleep(5)
	}
}
