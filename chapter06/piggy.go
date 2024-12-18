package main

import (
	"fmt"
	"math/rand"
)

func main() {
	singleMoney := []float64{0.05, 0.10, 0.25}
	piggyBank := 0.0
	for piggyBank < 20.00 {
		piggyBank += singleMoney[rand.Intn(3)]
		fmt.Printf("piggy: %05.2f\n", piggyBank)
	}
}
