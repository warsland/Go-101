package main

import "fmt"

func main() { //匿名函数的声明和调用整合在同一步骤中
	func() { //声明匿名函数
		fmt.Println("Functions anonymous")
	}() //调用匿名函数
}
