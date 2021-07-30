package marionette

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func GoAFK() {
	robotgo.MoveMouseSmooth(341, 65)
	robotgo.MouseClick("left", false)
	time.Sleep(RandDuration(500, 800))

	robotgo.MoveMouseSmooth(928, 662)
	robotgo.MouseClick("left", false)
	time.Sleep(RandDuration(500, 800))

	robotgo.MoveMouseSmooth(1096, 342)
	robotgo.MouseClick("left", false)
}
