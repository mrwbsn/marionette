package main

import (
	"log"

	"github.com/mrwbsn/marionette"
	"gocv.io/x/gocv"
)

func main() {
	wi := gocv.NewWindow("normal")
	wt := gocv.NewWindow("threshold")
	wt.ResizeWindow(600, 600)
	wt.MoveWindow(1528, 128)
	wi.MoveWindow(1528, 728)
	wi.ResizeWindow(600, 600)

	lh := wi.CreateTrackbar("Low H", 360/2)
	hh := wi.CreateTrackbar("High H", 255)
	ls := wi.CreateTrackbar("Low S", 255)
	hs := wi.CreateTrackbar("High S", 255)
	lv := wi.CreateTrackbar("Low V", 255)
	hv := wi.CreateTrackbar("High V", 255)

	for {
		img, err := gocv.ImageToMatRGBA(marionette.CaptureMonitor0())
		if err != nil {
			log.Fatal(err)
		}
		gocv.CvtColor(img, &img, gocv.ColorRGBAToBGR)
		gocv.CvtColor(img, &img, gocv.ColorBGRToHSV)
		thresholded := gocv.NewMat()
		gocv.InRangeWithScalar(img,
			gocv.Scalar{Val1: getPosFloat(lh), Val2: getPosFloat(ls), Val3: getPosFloat(lv)},
			gocv.Scalar{Val1: getPosFloat(hh), Val2: getPosFloat(hs), Val3: getPosFloat(hv)},
			&thresholded)

		wi.IMShow(img)
		wt.IMShow(thresholded)
		if wi.WaitKey(1) == 27 || wt.WaitKey(1) == 27 {
			break
		}
	}
}

func getPosFloat(t *gocv.Trackbar) float64 {
	return float64(t.GetPos())
}
