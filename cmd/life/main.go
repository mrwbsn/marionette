package main

import (
	"log"
	"time"

	"github.com/mrwbsn/marionette"
)

func main() {
	log.Println("bot start!")
	time.Sleep(3 * time.Second)

	activity := 1
	switch activity {
	case 0:
		marionette.GoGardening()
	case 1:
		marionette.GoFishing()
	}
}
