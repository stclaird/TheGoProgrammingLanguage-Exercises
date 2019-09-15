package conversion

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Pound float64
type Kilo float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (lb Pound) String() string     { return fmt.Sprintf("%gLB", lb) }
func (kg Kilo) String() string      { return fmt.Sprintf("%gKG", kg) }
