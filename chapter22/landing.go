package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	D, M, S float64
	H       rune
}

type location struct {
	Lat, Long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	ColumbibaMemorialSattion := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	ChallengerMemorialSatation := newLocation(coordinate{1, 56, 46.3, 'N'}, coordinate{354, 28, 24.2, 'E'})
	BradburyLanding := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	Elysium := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})
	locations := []location{ColumbibaMemorialSattion, ChallengerMemorialSatation, BradburyLanding, Elysium}
	for _, location := range locations {
		bytes, err := json.MarshalIndent(location, " ", "\t")
		exitOnError(err)
		fmt.Println(string(bytes))
	}
}
