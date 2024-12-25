package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32
	return c
}

func kelvinToFahrenheit(k float64) float64 {
	k -= 459.67
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "K is ", celsius, "℃\n")
	kelvin = 233.0
	celsius = kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "K is ", celsius, "℃\n")
	celsius = 37.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Print(celsius, "℃ is ", fahrenheit, "℉\n")
	kelvin = 0.0
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Print(kelvin, "K is ", fahrenheit, "℉\n")
}
