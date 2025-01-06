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

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (r report) average() celsius { //编写一个转发至实际实现的方法
	return r.temperature.average()
}

func main() {

	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %v℃\n", t.average())
	report := report{sol: 15, temperature: t}
	fmt.Printf("average %v℃\n", report.temperature.average())
	fmt.Printf("average %v℃\n", report.average())
}
