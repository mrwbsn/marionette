package main

import (
	"flag"
	"log"
	"time"

	"github.com/mrwbsn/marionette"
)

func main() {
	var activity int
	flag.IntVar(&activity, "act", 0, "0 = Gardening | 1 = Fishing")
	flag.Parse()

	time.Sleep(3 * time.Second)

	log.Println("bot start!")

	switch activity {
	case 0:
		marionette.GoGardening()
	case 1:
		marionette.GoFishing()
	}
}
