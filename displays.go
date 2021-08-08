package marionette

import "strconv"

type Display struct {
	Label string
	X     string
	Y     string
	Color string
}

func FindDisplay(key int) *Display {
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
