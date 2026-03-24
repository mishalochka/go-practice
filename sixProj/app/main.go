package main

import (
	"fmt"
	"os"
	"sixProj/tempconv"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fu := tempconv.Funt(t)
		fe := tempconv.Feet(t)
		kg := tempconv.Kilogramm(t)
		m := tempconv.Meter(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c),
			"%s = %s, %s = %s\n",
			k, tempconv.KToC(k), c, tempconv.CToK(c),
			"%s = %s, %s = %s\n",
			fu, tempconv.FuToKg(fu), kg, tempconv.KgToFu(kg),
			"%s = %s, %s = %s\n",
			fe, tempconv.FeToM(fe), m, tempconv.MToFe(m))

	}
}
