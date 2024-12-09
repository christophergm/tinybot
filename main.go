package main

import (
	"machine" ///grandcentral-m4"
	"time"
	// "tinygo.org/x/drivers/ws2812"
)

func main() {
	// n := ws2812.New(machine.PC24)

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

	// Configure the pin
	buttonB := machine.PC17
	buttonB.Configure(machine.PinConfig{Mode: machine.PinTimer})

	buttonR := machine.PC16
	buttonR.Configure(machine.PinConfig{Mode: machine.PinTimer})

	// Set up PWM timer
	pwmTimer := machine.TCC0
	pwmTimer.Configure(machine.PWMConfig{})

	chB, err := pwmTimer.Channel(buttonB)
	if err != nil {
		println("Failed to get PWM channel:", err)
		return
	}

	chR, err := pwmTimer.Channel(buttonR)
	if err != nil {
		println("Failed to get PWM channel:", err)
		return
	}

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
			pwmTimer.Set(chB, i)
			time.Sleep(time.Millisecond * 25)
		}
	}()

	select {}
}
