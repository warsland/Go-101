package main

import "fmt"

var f = func() { //将匿名函数赋值给变量
	fmt.Println("Dress up for the masquerade")
}

func main() {
	f() //匿名函数的调用
}
