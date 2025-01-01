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

	pauseMilliseconds := 100

	rand.Seed(uint64(time.Now().UnixNano()))

	neoPixel := peripheral.NeoPixel{}
	neoPixel.Configure()

	boardYellowLight := peripheral.BoardYellowLight{}
	boardYellowLight.Configure()
	boardYellowLight.StartBlink()

	neoPixel.SetColorAndPause(Green, pauseMilliseconds)

	elevatorButton := peripheral.Elevator{Period: 1000}
	elevatorButton.Configure()
	go elevatorButton.Run()

	neoPixel.SetColorAndPause(Red, pauseMilliseconds)
	neoPixel.SetColorAndPause(Yellow, pauseMilliseconds)
	neoPixel.SetColorAndPause(Blue, pauseMilliseconds)

	// spi := peripheral.Spi{}
	// if err := spi.Configure(); err != nil {
	// 	for {
	// 		neoPixel.SetRandomColorAndPause(pauseMilliseconds)
	// 	}
	// }
	// spi.Start()

	dotStar := peripheral.DotStarRGB{}
	if err := dotStar.Configure(); err != nil {
		for {
			neoPixel.SetRandomColorAndPause(pauseMilliseconds)
		}
	}
	dotStar.Start()

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
