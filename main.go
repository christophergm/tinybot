package main

import (
	"image/color"
	"machine"
	"time"

	"golang.org/x/exp/rand"
	// "tinygo.org/x/drivers/apa102"
	"tinygo.org/x/drivers/ws2812"
)

func main() {

	rand.Seed(uint64(time.Now().UnixNano()))

	// Configure the onboard NeoPixel
	neoPixelPin := machine.PC24
	neoPixelPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neoPixelDriver := ws2812.NewWS2812(neoPixelPin)

	// // probably best to init serial for debugging if you need it
	//machine.InitSerial()

	// Initialize SPI and LED strip driver
	// spi := machine.SPI0 // all mosi and sck pins are pd09 and pd08
	// spi.Configure(machine.SPIConfig{
	// 	Frequency: 4000000,      // 4 MHz, typical for APA102
	// 	SCK:       machine.PD09, // SCK
	// 	SDO:       machine.PD08, // MOSI
	// })
	// SetRandomColor sets the NeoPixel to a random color

	time.Sleep(500 * time.Millisecond)

	neoPixelDriver.WriteColors([]color.RGBA{{255, 0, 0, 20}})

	SetRandomColor := func(led ws2812.Device) {
		// Generate random RGB values
		r := uint8(rand.Intn(256))
		g := uint8(rand.Intn(256))
		b := uint8(rand.Intn(256))

		// Write the color to the NeoPixel
		neoPixelDriver.WriteColors([]color.RGBA{{r, g, b, 20}})
	}

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
	// Configure the pins
	buttonLedB := machine.PC17
	buttonLedB.Configure(machine.PinConfig{Mode: machine.PinTimer})

	buttonLedR := machine.PC16
	buttonLedR.Configure(machine.PinConfig{Mode: machine.PinTimer})

	buttonInput := machine.PB13
	buttonInput.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	// Set up PWM timer
	pwmTimer := machine.TCC0
	pwmTimer.Configure(machine.PWMConfig{})

	chB, err := pwmTimer.Channel(buttonLedB)
	if err != nil {
		println("Failed to get PWM channel:", err)
		return
	}

	chR, err := pwmTimer.Channel(buttonLedR)
	if err != nil {
		println("Failed to get PWM channel:", err)
		return
	}

	// Initialize SPI and LED strip driver
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		Frequency: 4000000,      // 4 MHz, typical for APA102
		SCK:       machine.PD09, // SCK
		SDO:       machine.PD08, // MOSI
	})

	//ledStrip := apa102.New(spi)

	// go testSPI(ledStrip)

	// numLEDs := 144 // Number of LEDs on your strip
	// buffer := make([]color.RGBA, numLEDs)

	// Set the PWM period (frequency)
	pwmTimer.SetPeriod(1000) // 1 kHz

	// Ramp up and down the PWM duty cycle
	max := pwmTimer.Top()
	go func() {
		i := uint32(0)
		onCount := max / uint32(10)
		direction := 1
		pwmTimer.Set(chR, max/10)
		for {
			if direction == 1 {
				i = i + onCount
			} else {
				i = i - onCount
			}
			if i >= max {
				i = max
				direction = -1
			}
			if i <= 0 {
				i = 0
				direction = 1
			}
			if buttonInput.Get() == true {
				pwmTimer.Set(chR, 0)
			} else {
				pwmTimer.Set(chR, max/10)
			}
			pwmTimer.Set(chB, i)
			time.Sleep(time.Millisecond * 25)
		}
	}()

	// Continuously set a random color every second
	for {
		SetRandomColor(neoPixelDriver)
		time.Sleep(1 * time.Second)
	}
	select {}
}
