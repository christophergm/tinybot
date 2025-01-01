package main

import (
	"image/color"
	"time"

	"golang.org/x/exp/rand"

	"github.com/chris/tinybot/peripheral"
	"tinygo.org/x/drivers/apa102"
)

var (
	Red    = color.RGBA{25, 0, 0, 20}
	Green  = color.RGBA{0, 25, 0, 20}
	Blue   = color.RGBA{0, 0, 25, 20}
	Yellow = color.RGBA{25, 25, 0, 20}
)

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))

	neoPixel := peripheral.NeoPixel{}
	neoPixel.Configure()

	boardYellowLight := peripheral.BoardYellowLight{}
	boardYellowLight.Configure()
	boardYellowLight.StartBlink()

	neoPixel.SetColorAndPause(Green, 500)

	// elevatorButton := peripheral.Elevator{Period: 1000}
	// go elevatorButton.Run()

	neoPixel.SetColorAndPause(Red, 500)
	neoPixel.SetColorAndPause(Yellow, 500)
	neoPixel.SetColorAndPause(Blue, 500)

	// Continuously set a random color every second
	for {
		neoPixel.SetRandomColorAndPause(500)
		time.Sleep(1 * time.Second)
	}

	select {}
}

func testSPI(ledStrip *apa102.Device) {
	for {
		ledStrip.WriteColors([]color.RGBA{
			{R: 255, G: 0, B: 0, A: 255},
			{R: 0, G: 255, B: 0, A: 255},
			{R: 0, G: 0, B: 255, A: 255},
		})
		time.Sleep(500 * time.Millisecond)
	}
}
