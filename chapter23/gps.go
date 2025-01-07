package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

type gps struct {
	w         world
	now, dest location
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (l location) descripition() string {
	return fmt.Sprintf("Name: %v,Latitude: %v,Longitude: %v", l.name, l.lat, l.long)
}

func (g gps) distance() float64 {
	return g.w.distance(g.now, g.dest)
}

func (g gps) message() {
	dist := fmt.Sprintf("the distance to destination is %v km\n", g.distance())
	fmt.Printf("now we are at %v ,the destination is %v,%v", g.now.descripition(), g.dest.descripition(), dist)
}

type rover struct {
	gps
}

func main() {
	curiosity := rover{
		gps{
			w:    world{radius: 3389.5},
			now:  location{name: "Bradbury", lat: -4.5895, long: 137.4417},
			dest: location{name: "Elysium", lat: 4.5, long: 135.9},
		},
	}

	curiosity.message()
}
