package main

import "fmt"

type fahrenheit float64
type celsius float64

type converter func(c celsius) fahrenheit

var f = func() {
	fmt.Println("=================")
}

func celsiusToFahrenhrit(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func drawTable(con converter, offset celsius) fahrenheit {
	return con(offset)
}

func main() {
	var degree celsius = -40.0
	f()
	fmt.Printf("|  ℃   |  ℉   |\n")
	for i := degree; i <= 100.0; i += 5 {
		f()
		d := drawTable(celsiusToFahrenhrit, i)
		fmt.Printf("|%-6.2f|%-6.2f|\n", i, d)
	}
	f()

}
