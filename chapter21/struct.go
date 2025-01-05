package main

import "fmt"

func main() {
	var curisity struct {
		lat  float64
		long float64
	}

	curisity.lat = -4.5895
	curisity.long = 137.4417

	fmt.Println(curisity.lat, curisity.long)
	fmt.Println(curisity)
}
