package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	D, M, S float64
	H       rune
}

type location struct {
	Lat, Long float64
}

type world struct {
	Radius float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))

	return w.Radius * math.Acos(s1*s2+c1*c2*clong)
}

func disted(name string, dis float64, d map[string][]float64) {
	d[name] = append(d[name], dis)
}

func main() {
	Mars := world{3389.5}
	Earth := world{6371.0}

	ColumbibaMemorialSattion := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	ChallengerMemorialSatation := newLocation(coordinate{1, 56, 46.3, 'N'}, coordinate{354, 28, 24.2, 'E'})
	BradburyLanding := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	Elysium := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})
	dist := make(map[string][]float64, 6)
	disted("cToCh", Mars.distance(ColumbibaMemorialSattion, ChallengerMemorialSatation), dist)
	disted("cToB", Mars.distance(ColumbibaMemorialSattion, BradburyLanding), dist)
	disted("cToE", Mars.distance(ColumbibaMemorialSattion, Elysium), dist)
	disted("chToB", Mars.distance(ChallengerMemorialSatation, BradburyLanding), dist)
	disted("chToE", Mars.distance(ChallengerMemorialSatation, Elysium), dist)
	disted("bToE", Mars.distance(BradburyLanding, Elysium), dist)
	for n, d := range dist {
		fmt.Printf("%v\t%v\n", n, d)
	}

	fmt.Printf("London to Pairs is %v km\n", Earth.distance(newLocation(coordinate{51, 30, 0.0, 'N'}, coordinate{0, 8, 0.0, 'W'}), newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})))
	fmt.Printf("distance is %v km\n", Mars.distance(newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0.0, 'E'}), newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})))
}
