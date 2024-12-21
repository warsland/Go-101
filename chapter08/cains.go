package main

import (
	"fmt"
	"math/big"
)

func main() {
	distance := new(big.Int)
	distance.SetString("236000000000000000", 10)
	lightSpeed := big.NewInt(299792)
	daysPerDays := big.NewInt(365)
	secondsPerDays := big.NewInt(86400)
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDays)
	cains := new(big.Int)
	cains.Div(days, daysPerDays)
	fmt.Println("The Cains Major Dwarf is", cains, "ly wawy")
}
