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

type ddAndDms struct {
	Decimal   float64 `json:"decimal"`
	Dms       string  `json:"dms"`
	Degrees   float64 `json":"degrees`
	Minutes   float64 `json":"minutes`
	Seconds   float64 `json":"seconds`
	Hemispher rune    `json":"hemispher`
}

type location struct {
	Lat, Long coordinate
}

type Marshaler interface {
	Marshal()
}

func (l location) String() string {
	return fmt.Sprintf("%v %v", l.Lat, l.Long)
}

func (c coordinate) decimal() float64 {
	sgin := 1.0
	switch c.H {
	case 's', 'w', 'S', 'W':
		sgin = -1.0
	}

	return sgin * (c.D + c.M/60 + c.S/3600)
}

func (c coordinate) Marshaler() {
	var ddd = ddAndDms{
		Decimal:   c.decimal(),
		Dms:       fmt.Sprintf("%v%v%v%v", c.D, c.M, c.S, c.H),
		Degrees:   c.D,
		Minutes:   c.M,
		Seconds:   c.S,
		Hemispher: c.H,
	}
	bytes, err := json.MarshalIndent(ddd, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}

func main() {
	elysim := location{
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia si at", elysim)
	elysim.Lat.Marshaler()
}
