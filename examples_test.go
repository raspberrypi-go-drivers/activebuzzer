package activebuzzer_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/activebuzzer"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	// buzzer connected to GPIO pin 17
	buzzer := activebuzzer.NewActiveBuzzer(17)
	buzzer.Tone()
	time.Sleep(3 * time.Second)
	buzzer.StopTone()
}

func ExampleNewActiveBuzzer() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	// buzzer connected to GPIO pin 17
	buzzer := activebuzzer.NewActiveBuzzer(17)
	buzzer.Tone()
}

func ExampleActiveBuzzer_Tone() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	// buzzer connected to GPIO pin 17
	buzzer := activebuzzer.NewActiveBuzzer(17)
	buzzer.Tone()
}

func ExampleActiveBuzzer_StopTone() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	// buzzer connected to GPIO pin 17
	buzzer := activebuzzer.NewActiveBuzzer(17)
	buzzer.Tone()
	buzzer.StopTone()
}

func ExampleActiveBuzzer_ToggleTone() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("impossible to init gpio")
		os.Exit(1)
	}
	defer rpio.Close()
	// buzzer connected to GPIO pin 17
	buzzer := activebuzzer.NewActiveBuzzer(17)
	for i := 0; i < 10; i++ {
		buzzer.ToggleTone()
	}
}
