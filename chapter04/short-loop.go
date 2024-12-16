package main

import "fmt"

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	} //循环结束，count变量不在处于作用域内。
}
