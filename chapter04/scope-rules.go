package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" //era变量在整个包内都是可用的

func main() {
	year := 2018                               //era变量和year变量都处于main函数作用域内
	switch month := rand.Intn(12) + 1; month { //era变量、year变量和month变量都处于switch作用域内
	case 2: //era变量、year变量、month变量和day变量都处于该case作用域内
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11: //该case作用域内的day变量和上一个case作用域内的day变量不是同一个变量。
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	case 1, 3, 5, 7, 8, 10, 12:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
