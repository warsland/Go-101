package main

import (
	"fmt"
	"math"
)

type report struct {
	sol
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
type sol int

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int { //基于sol声明的方法可以通过sol或者report进行访问
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}
