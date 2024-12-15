package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = 55
	for {
		var guessnum = rand.Intn(100)
		if num > guessnum {
			fmt.Printf("The guess num is %v,and it's bigger\n", guessnum)
		} else if num < guessnum {
			fmt.Printf("The guess num is %v,and it's smaller\n", guessnum)
		} else {
			fmt.Printf("The guess num is %v,and it's right\n", guessnum)
			break
		}
	}
}
