package main

import "fmt"

func main() {
	var soup map[string]int
	fmt.Println(soup == nil)

	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}
	for ingredient, memeasurement := range soup {
		fmt.Println(ingredient, memeasurement)
	}

}
