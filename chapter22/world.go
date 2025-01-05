package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type loaction struct {
	lat, long float64
}

func rad(deg float64) float64 { //将角度转化为弧度
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 loaction) float64 { //使用宇轩球面定律计算两个位置之间的关系
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)

}

func main() {
	var mars = world{radius: 3389.5}

	spirit := loaction{-14.5684, 175.472636}
	opportunity := loaction{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
}
