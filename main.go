package main

import (
	"machine" ///grandcentral-m4"
	"time"
	// "tinygo.org/x/drivers/ws2812"
)

func main() {
	led := machine.PC30
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure the pin
	buttonRed := machine.PC17
	buttonRed.Configure(machine.PinConfig{Mode: machine.PinTimer})

	// Set up PWM timer for ping
	pwmTimer := machine.TCC0
	pwmTimer.Configure(machine.PWMConfig{})

	ch, err := pwmTimer.Channel(buttonRed)
	if err != nil {
		println("Failed to get PWM channel:", err)
		return
	}

	// Set the PWM period (frequency)
	pwmTimer.SetPeriod(1000) // 1 kHz

	// n := ws2812.New(machine.PC24)

	// Blink yellow board LED
	go func() {
		for {
			led.Low()
			buttonRed.Low()
			time.Sleep(time.Millisecond * 250)

			led.High()
			buttonRed.High()
			time.Sleep(time.Millisecond * 250)
		}
	}()

	// Ramp up and down the PWM duty cycle
	go func() {
		i := uint32(0)
		onCount := pwmTimer.Top() / uint32(30)
		direction := 1
		for {
			if direction == 1 {
				i = i + onCount
			} else {
				i = i - onCount
			}
			if i >= pwmTimer.Top() {
				i = pwmTimer.Top()
				direction = -1
			}
			if i <= 0 {
				i = 0
				direction = 1
			}
			pwmTimer.Set(ch, i)
			time.Sleep(time.Millisecond * 25)
		}
	}()

	select {}
}
