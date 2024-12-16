package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 { //开启新的作用域
		num := rand.Intn(10) + 1 //变量的简短声明，与var num = rand.Intn(10) + 1 相同
		fmt.Println(num)
		count++
	} //作用域结束
}
