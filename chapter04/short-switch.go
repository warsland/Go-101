package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(3); num {
	case 0:
		fmt.Println("Speace Adventures")
	case 1:
		fmt.Println("SpeaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}
