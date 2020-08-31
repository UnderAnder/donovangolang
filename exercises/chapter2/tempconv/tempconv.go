package tempconv

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprint("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprint("%g°F", f)
}
