# Active Buzzer

[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/activebuzzer.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/activebuzzer)
![golangci-lint](https://github.com/raspberrypi-go-drivers/activebuzzer/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/activebuzzer)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/activebuzzer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Package activebuzzer is a driver allowing to control an active buzzer from GPIO pin

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/activebuzzer.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/activebuzzer)

## Quick start

```go
import (
	"fmt"

	"github.com/raspberrypi-go-drivers/button"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
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
```

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH using integrated bluetooth but may work well on other Raspberry Pi having integrated Bluetooth

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)
