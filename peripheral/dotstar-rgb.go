package peripheral

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/apa102"
)

type DotStarRGB struct {
	buffer []color.RGBA
	// Number of LEDs on your strip
	numLEDs  int
	ledStrip *apa102.Device
	startPos int
	canceled bool
}

func (d *DotStarRGB) Configure() error {

	d.numLEDs = 144
	d.buffer = make([]color.RGBA, d.numLEDs)

	// Initialize SPI and LED strip driver
	spi := machine.SPI0
	err := spi.Configure(machine.SPIConfig{
		// Frequency: 4000000,      // 4 MHz, typical for APA102
		// SCK:       machine.PD09, // SCK
		// SDO:       machine.PD08,
		// SDI:       machine.PA22,
		// Mode:      0, // MOSI
	})
	if err != nil {
		return err
	}

	d.ledStrip = apa102.New(spi)
	return nil
}

func (d *DotStarRGB) StartSpin() {
	var col color.RGBA
	d.canceled = false
	go func() {
		tailLength := rand.Intn(40) + 1
		for {

			for i := 0; i < d.numLEDs; i++ {

				if i < tailLength {
					col = color.RGBA{R: uint8(i * 3), G: uint8(i / 2), B: uint8(i), A: 255}
				} else {
					col = color.RGBA{R: 0, G: 0, B: 0, A: 255}
				}

				pos := (i + d.startPos) % d.numLEDs
				d.buffer[pos] = col
			}

			d.ledStrip.WriteColors(d.buffer)

			time.Sleep(20 * time.Millisecond)
			d.startPos++
			if d.canceled {
				return
			}
		}
	}()
}

func (d *DotStarRGB) Explode() {
	d.canceled = true
	time.Sleep(50 * time.Millisecond)

	for j := 0; j < 10; j++ {

		for i := 0; i < d.numLEDs; i++ {
			distance := (d.startPos - i) % d.numLEDs
			magnitude := 3 * (d.numLEDs - distance) / d.numLEDs
			magnitude = magnitude + rand.Intn(9) - j
			if magnitude < 0 {
				magnitude = 0
			}
			col := color.RGBA{R: uint8(3 * magnitude), G: uint8(2 * magnitude), B: uint8(magnitude), A: 255}
			d.buffer[i] = col
		}

		d.ledStrip.WriteColors(d.buffer)

		time.Sleep(20 * time.Millisecond)
	}
}
