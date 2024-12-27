package main

import "fmt"

func main() {
	f := func(message string) { //将匿名函数赋值给变量
		fmt.Println(message)
	}

	f("Go to the party")
}
