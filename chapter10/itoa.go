package main

import (
	"fmt"
	"strconv"
)

func main() {
	cooutdown := 10
	str := "Launch in T minus " + strconv.Itoa(cooutdown) + " seconds."
	fmt.Println(str)
}
