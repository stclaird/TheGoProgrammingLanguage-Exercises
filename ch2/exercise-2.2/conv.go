package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/conversion"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conversion.Fahrenheit(t)
		c := conversion.Celsius(t)
		lb := conversion.Pound(t)
		kg := conversion.Kilo(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			f, conversion.FToC(f),
			f, conversion.FToK(f),
			c, conversion.CToF(c),
			c, conversion.CToK(c),
			lb, conversion.PToK(lb),
			kg, conversion.KToP(kg))
	}
}
