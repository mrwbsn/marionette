package marionette

import (
	"fmt"
	"math/rand"
	"time"

	"gocv.io/x/gocv"
)

func RandDuration(min, max int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(max-min+1) + min
	dur, _ := time.ParseDuration(fmt.Sprintf("%vs", randInt))
	return dur
}

func DebugCV(img gocv.Mat) {
	win := gocv.NewWindow("Debug")
	for {
		win.IMShow(img)
		if win.WaitKey(1) == int('q') {
			win.Close()
			break
		}
	}
}
