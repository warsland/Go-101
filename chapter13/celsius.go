package main

import "fmt"

func main() {
	type celsius float64 //新类型名为celsius,底层类型为float64
	var temperature celsius = 20
	fmt.Println(temperature)
}
