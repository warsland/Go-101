package main

import "fmt"

type fahrenheit float64
type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

var line = func() {
	fmt.Println("=======================")
}

func celsiusToFahrenheit(degree float64) (string, string) {
	celsius := celsius(degree)
	fahrenheit := celsius.fahrenheit()
	celsius_1 := fmt.Sprintf("%.1f", celsius)
	fahrenheit_1 := fmt.Sprintf("%1.f", fahrenheit)
	return celsius_1, fahrenheit_1
}

func fahrenheitToCelsius(degree float64) (string, string) {
	fahrenheit := fahrenheit(degree)
	celsius := fahrenheit.celsius()
	fahrenheit_1 := fmt.Sprintf("%.1f", fahrenheit)
	celsius_1 := fmt.Sprintf("%.1f", celsius)
	return fahrenheit_1, celsius_1
}

func drawTable(hdr1, hdr2 string, degree float64, getRow func(degree float64) (string, string)) {
	line()
	fmt.Printf("| %-8s | %-8s |\n", hdr1, hdr2)
	for i := -40.0; i <= degree; i += 5.0 {
		line()
		number_string1, number_string2 := getRow(i)
		fmt.Printf("| %-8s | %-8s |\n", number_string1, number_string2)
	}
	line()
}

func main() {
	drawTable("℃", "℉", 100, celsiusToFahrenheit)
	fmt.Println()
	drawTable("℉", "℃", 100, fahrenheitToCelsius)
}
