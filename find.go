package marionette

import (
	"image"
	"log"

	"gocv.io/x/gocv"
)

type MatchTemplate struct {
	Threshold float32
}

func (mt MatchTemplate) Find(path string) (image.Point, float32) {
	var location image.Point
	var confidence float32

	findOutput := gocv.NewMat()
	defer findOutput.Close()

	mask := gocv.NewMat()
	defer mask.Close()

	screenThres := gocv.NewMat()
	defer screenThres.Close()

	templateThres := gocv.NewMat()
	defer templateThres.Close()

	for {
		screen, err := gocv.ImageToMatRGBA(CaptureMonitor0())
		if err != nil {
			log.Fatal(err)
		}

		template := gocv.IMRead(path, gocv.IMReadFlag(gocv.MatTypeCV8UC4))

		gocv.MatchTemplate(screen, template, &findOutput, gocv.TmCcoeffNormed, mask)

		_, maxVal, _, maxLoc := gocv.MinMaxLoc(findOutput)

		if maxVal > mt.Threshold {
			location = maxLoc
			confidence = maxVal
			break
		}

		log.Println("waiting..")

	}

	return location, confidence
}
