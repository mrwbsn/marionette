package marionette

import (
	"encoding/json"

	"github.com/go-vgo/robotgo"
)

type Displays struct {
	Auto struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"auto"`
	Life struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"life"`
	Gardening struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gardening"`
	ShatteringShroom struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"shattering_shroom"`
	GardeningPlaceC struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gardening_place_c"`
	ExitWindow struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"exit_window"`
}

func NewDisplayInfo() Displays {
	var displays Displays
	json.Unmarshal(ReadJSON("./configs/displays.json"), &displays)
	return displays
}

func clickDisplay(x, y int) {
	robotgo.MoveMouse(x, y)
	robotgo.MouseClick("left", false)
	robotgo.Sleep(1)
}

func (d *Displays) ClickAuto() {
	clickDisplay(d.Auto.X, d.Auto.Y)
}

func (d *Displays) ClickLife() {
	clickDisplay(d.Life.X, d.Life.Y)
}

func (d *Displays) ClickGardening() {
	clickDisplay(d.Gardening.X, d.Gardening.Y)
}

func (d *Displays) ClickShatteringShroom() {
	// TODO: need to refactor, so it can be use for all plants
	clickDisplay(d.ShatteringShroom.X, d.ShatteringShroom.Y)
}

func (d *Displays) ClickGardeningPlaceC() {
	clickDisplay(d.GardeningPlaceC.X, d.GardeningPlaceC.Y)
}

func (d *Displays) ClickExitWindow() {
	clickDisplay(d.ExitWindow.X, d.ExitWindow.Y)
}
