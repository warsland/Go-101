package main

import "fmt"

func main() {
	age := 41
	marAge := float64(age)

	marsDay := 687.0
	erthDays := 365.2425
	marAge = marAge * erthDays / marsDay
	fmt.Println("I am", marAge, "years old on Mars.")
}
