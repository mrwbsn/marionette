package marionette

import (
	"image"
	"log"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
)

// CaptureMonitor0 captures main monitor
func CaptureMonitor0() image.Image {
	bounds := GetDisplayBound()
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Fatalln(err)
	}

	return img
}

func CaptureFromBitmap() image.Image {
	return robotgo.ToImage(robotgo.ToCBitmap(robotgo.GoCaptureScreen()))
}

func GetDisplayBound() image.Rectangle {
	return screenshot.GetDisplayBounds(0)
}

func GetScreenSize() image.Point {
	bounds := GetDisplayBound()
	return bounds.Size()
}
