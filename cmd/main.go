package main

import (
	"flag"
	"log"
	"time"

	"github.com/mrwbsn/marionette"
)

// TODO: monitor scale
// const (
// 	h = 1440
// 	w = 900
// )

func main() {

	var activity int
	flag.IntVar(&activity, "gardening", 0, "0 = Gardening | 1 = Fishing")

	log.Printf("screen size: %v", marionette.GetScreenSize())

	time.Sleep(3 * time.Second)

	log.Println("bot start!")

	switch activity {
	case 0:
		marionette.GoGardening()
	case 1:
		marionette.GoFishing()
	}

}
