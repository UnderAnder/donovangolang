package main

import "fmt"

type Metre float64
type Foot float64

func (m Metre) String() string {
	return fmt.Sprintf("%g метров", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g футов", f)
}

func MToF(m Metre) Foot {
	return Foot(m * 3.281)
}

func FToM(f Foot) Metre {
	return Metre(f / 3.281)
}
