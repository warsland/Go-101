// 浮点数不精准的例子
package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third + third + third)

	piggbyBank := 0.1
	piggbyBank += 0.2
	fmt.Println(piggbyBank)
}
