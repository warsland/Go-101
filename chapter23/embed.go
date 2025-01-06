package main

import "fmt"

type report struct {
	sol         int
	temperature //将temperature类型嵌入report中
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Printf("acerage %v ℃ \n", report.average())

	fmt.Printf("acerage %v ℃ \n", report.temperature.average())

	fmt.Printf(" %v ℃ \n", report.high)
	report.high = 32
	fmt.Printf(" %v ℃ \n", report.temperature.high)

}
