package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func MToFe(m Meter) Feet { return Feet(m / OneFeet) }

func FeToM(fe Feet) Meter { return Meter(fe * OneFeet) }

func FuToKg(fu Funt) Kilogramm { return Kilogramm(fu / OneFunt) }

func KgToFu(kg Kilogramm) Funt { return Funt(kg * OneFunt) }
