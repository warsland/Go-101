package main

import "fmt"

func main() {
	year := 2024
	fmt.Printf("Type %T for %v\n", year, year)
	fmt.Printf("Type %T for %[1]v\n", year)
	//在格式化变量%v前添加[1]来复用前一个格式化变量的值，避免一个变量在Printf函数中重复使用
}
