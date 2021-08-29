package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	// for i := 0; i < 200; i++ {
	// 	fmt.Println(GetCorByCoor(1125, 525))
	// }
	fmt.Println(GetColorByCoor(1175, 20))
}

func GetColorByCoor(x, y int) string {
	c := robotgo.GetPixelColor(x, y)
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}

func GetColorByMouse() string {
	c := robotgo.GetMouseColor()
	x, y := robotgo.GetMousePos()
	return fmt.Sprintf("%v,%v,%v", x, y, c)
}
