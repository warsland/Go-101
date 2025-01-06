package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	loaction    loaction
}

type temperature struct {
	high, low celsius
}

type loaction struct {
	lat, long float64
}

type celsius float64

func main() {

	bradbury := loaction{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, loaction: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v℃\n", report.temperature.high)
}
