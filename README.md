# Button

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/activebuzzer)

Package activebuzzer is a driver allowing to control an avtive buzzer from GPIO pin

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/activebuzzer)

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
