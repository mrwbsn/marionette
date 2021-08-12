package marionette

import (
	"log"
	"strconv"
)

type Display struct {
	Label string
	X     string
	Y     string
	Color string
}

func Find(key int) *Display {
	displays := ReadCSV("./configs/displays.csv")
	for _, display := range displays {
		if display[0] == strconv.Itoa(key) {
			return &Display{
				Label: display[1],
				X:     display[2],
				Y:     display[3],
				Color: display[4],
			}
		}
	}
	return &Display{}
}

func Get(d *Display) (int, int, string) {
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
