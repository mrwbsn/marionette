package marionette

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"gocv.io/x/gocv"
)

func RandDuration(min, max int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(max-min+1) + min
	dur, _ := time.ParseDuration(fmt.Sprintf("%vms", randInt))
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

func ReadJSON(filePath string) []byte {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	jsonVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	return jsonVal
}

func ReadCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	r, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return r
}
