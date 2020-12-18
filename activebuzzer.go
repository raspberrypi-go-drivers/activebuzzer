package activebuzzer

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// ActiveBuzzer represent an active buzzer
type ActiveBuzzer struct {
	pin rpio.Pin
}

// NewActiveBuzzer creates a new ActiveBuzzer instance
func NewActiveBuzzer(pinID int) *ActiveBuzzer {
	buzzer := ActiveBuzzer{
		pin: rpio.Pin(uint8(pinID)),
	}
	buzzer.pin.Mode(rpio.Output)
	return &buzzer
}

// Tone activate the buzzer tone
func (buzzer *ActiveBuzzer) Tone() {
	buzzer.pin.High()
}

// StopTone stops the buzzer tone
func (buzzer *ActiveBuzzer) StopTone() {
	buzzer.pin.Low()
}

// ToggleTone toggle the buzzer tone
// Tone if stopped and stopped is toned
func (buzzer *ActiveBuzzer) ToggleTone() {
	if buzzer.pin.Read() == 1 {
		buzzer.pin.Low()
	} else {
		buzzer.pin.High()
	}
}
