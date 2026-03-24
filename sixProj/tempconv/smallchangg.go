package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Kilogramm float64
type Funt float64
type Meter float64
type Feet float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	OneFeet               = 0.3048
	OneFunt               = 2.2046
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (fu Funt) String() string      { return fmt.Sprintf("%g°K", fu) }
func (m Meter) String() string      { return fmt.Sprintf("%g°K", m) }
func (fe Feet) String() string      { return fmt.Sprintf("%g°K", fe) }
