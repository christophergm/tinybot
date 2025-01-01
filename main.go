package main

import (
	"image/color"
	"machine"
	"time"

	"golang.org/x/exp/rand"

	"github.com/chris/tinybot/peripheral"
	"tinygo.org/x/drivers/apa102"
)

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))

	neoPixel := peripheral.NeoPixel{}
	neoPixel.Configure()

	// Blink yellow board LED
	led := machine.PC30

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 250)

			led.High()
			time.Sleep(time.Millisecond * 250)
		}
	}()

	red := color.RGBA{255, 0, 0, 20}
	blue := color.RGBA{0, 255, 0, 20}
	green := color.RGBA{0, 0, 255, 20}
	yellow := color.RGBA{255, 255, 0, 20}

	neoPixel.SetColor(green)

	// Initialize SPI and LED strip driver
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		Frequency: 4000000,      // 4 MHz, typical for APA102
		SCK:       machine.PD09, // SCK
		SDO:       machine.PD08, // MOSI
	})

	// spi.configurePin()

	ledStrip := apa102.New(spi)
	go testSPI(ledStrip)

	// numLEDs := 144 // Number of LEDs on your strip
	// buffer := make([]color.RGBA, numLEDs)
	elevatorButton := peripheral.Elevator{Period: 1000}
	go elevatorButton.Run()

	neoPixel.SetColor(yellow)
	neoPixel.SetColor(red)
	neoPixel.SetColor(blue)

	// Continuously set a random color every second
	for {
		neoPixel.SetRandomColor()
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
