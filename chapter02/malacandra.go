package main

import (
	"fmt"
)

func main() {
	const (
		distance = 56000000
		time     = 28 * 24
	)
	var speed = distance / time
	fmt.Println(speed)
}
