package main

import (
	"flag"
	"fmt"
	"github.com/Tiril14/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var C float64
var F float64
var K float64
var out string
nn

func init() {
	flag.Float64Var(&C, "C", 0.0, "temperatur i grader celsius")

	flag.Float64Var(&F, "F", 0.0, "temperatur i grader fahrenheit")

	flag.Float64Var(&K, "K", 0.0, "temperatur i grader kelvin")

	flag.StringVar(&out, "out", "", "temperaturskala for resultat")
}

func main() {
	flag.Parse()
	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%g°F is %.2f°C\n", F, conv.FahrenheitToCelsius(F))
	} else if out == "K" && isFlagPassed("F") {
		fmt.Printf("%g°F is %.2f°K\n", F, conv.FahrenheitToKelvin(F))
	} else if out == "F" && isFlagPassed("C") {
		fmt.Printf("%g°C is %.2f°F\n", C, conv.CelsiusToFahrenheit(C))
	} else if out == "K" && isFlagPassed("C") {
		fmt.Printf("%g°C is %.2f°K\n", C, conv.CelsiusToKelvin(C))
	} else if out == "C" && isFlagPassed("K") {
		fmt.Printf("%g°K is %.2f°C\n", K, conv.KelvinToCelsius(K))
	} else if out == "F" && isFlagPassed("K") {
		fmt.Printf("%g°K is %.2f°F\n", K, conv.KelvinToFahrenheit(K))
	} else {
		fmt.Println("error: one of the following temperature arguments must be specified: -F, -C, -K")
		return
	}
}
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found

}
