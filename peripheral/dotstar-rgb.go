package peripheral

import (
	"image/color"
	"machine"
)

type DotStarRGB struct {
	buffer []color.RGBA
	// Number of LEDs on your strip
	numLEDs int
}

func (d *DotStarRGB) Configure() {

	d.numLEDs = 144
	d.buffer = make([]color.RGBA, d.numLEDs)

	// Initialize SPI and LED strip driver
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		Frequency: 4000000,      // 4 MHz, typical for APA102
		SCK:       machine.PD09, // SCK
		SDO:       machine.PD08, // MOSI
	})
}

// spi.configurePin()

// ledStrip := apa102.New(spi)
// go testSPI(ledStrip)

// numLEDs := 144 // Number of LEDs on your strip
//
