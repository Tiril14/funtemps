package conv

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32.0) * (5.0 / 9.0)

}

// Konverterer CelciusToFahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

// Konverterer KelvinToFahrenheit
func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*(9/5) + 32

}

// Konverterer FahrenheitToKelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*(5/9) + 273.15
}

// Konverterer CelsiusToKelvin
func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}

// Konverterer KelvinToCelsius
func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}
