package main

import (
	"fmt"
	"math/rand"
)

var penny = []float64{5, 10, 15}

func main() {
	number := 0
	money := 0.0
	fmt.Println("数量  币值  总额")
	for money < 20.00 {
		pennies := penny[rand.Intn(3)] / 100
		money += pennies
		number++
		fmt.Printf("%3v   %.2f  %.2f\n", number, pennies, money)

	}
}
