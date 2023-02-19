package main

import (
	"flag"
	"fmt"
)

// Definerer flag-variablene i hoved-"scope"
var C float64
var F float64
var K float64
var out string


func init() {
	flag.Float64Var(&C
	"C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&F
	"F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&K
	"K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "", "temperaturskala for resultat")
}

func main() {
	flag.Parse()
	var res float64
	//standardverdi 0
	if C != 0 {
		if out == "F"
	}
	res = conv.CelsiusToFahrenheit(C)
    } else if out == "K" {
	res = C + 273.15
	}
   }
   	fmt.Printf("%.2f\n",res)
}