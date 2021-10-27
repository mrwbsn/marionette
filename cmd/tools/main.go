package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	// x, y := 1585, 725
	// robotgo.MoveMouse(x, y)
	// robotgo.Sleep(5)
	// robotgo.Click()

	for {
		fmt.Println(GetColorByCoor(900, 347))
	}

	// robotgo.MoveMouse()

}

func GetColorByCoor(x, y int) string {
	c := robotgo.GetPixelColor(x, y)
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}

func GetColorByMouse() string {
	// robotgo.Sleep(3)
	c := robotgo.GetMouseColor()
	x, y := robotgo.GetMousePos()
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}
