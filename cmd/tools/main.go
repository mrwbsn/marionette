package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	for i := 0; i < 200; i++ {
		fmt.Println(GetCorByCoor(1122, 519))
	}
	// fmt.Println(GetColorByMouse())
}

func GetCorByCoor(x, y int) string {
	robotgo.MoveMouse(x, y)
	c := robotgo.GetMouseColor()
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}

func GetColorByMouse() string {
	c := robotgo.GetMouseColor()
	x, y := robotgo.GetMousePos()
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}
