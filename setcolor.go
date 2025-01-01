package main

import (
	"image/color"
	"time"

	"golang.org/x/exp/rand"
	"tinygo.org/x/drivers/ws2812"
)

// SetRandomColor sets the NeoPixel to a random color
func SetRandomColor(led ws2812.Device) {
	// Generate random RGB values
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	// Write the color to the NeoPixel
	led.WriteColors([]color.RGBA{{r, g, b, 255}})
	time.Sleep(time.Millisecond * 500)
}

func SetColor(led ws2812.Device, col color.RGBA) {
	// Write the color to the NeoPixel
	led.WriteColors([]color.RGBA{{col.R, col.G, col.B, 20}})
	time.Sleep(time.Millisecond * 500)
}
