package main

import (
	"machine" // grandcentral-m4"
	"time"
	// "tinygo.org/x/drivers/ws2812"
)

func main() {
	// n := ws2812.New(machine.PC24)

	// probably best to init serial for debugging if you need it
	machine.InitSerial()

	// wait a half second so serial can sync
	time.Sleep(time.Millisecond * 500)

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
	buttonInput.Configure(machine.PinConfig{Mode: machine.PinInput})
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
				pwmTimer.Set(chR, max)

			}
			pwmTimer.Set(chB, i)
			time.Sleep(time.Millisecond * 25)
		}
	}()

	select {}
}
