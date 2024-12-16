package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Speace Adventures")
	} else if num == 1 {
		fmt.Println("SpeaceX")
	} else {
		fmt.Println("Virgin Galactic")
	} //if语句结束，num变量不在处于作用域中。
}
