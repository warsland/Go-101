package main

import "fmt"

func main() {
	var count = 0
	for count = 10; count > 0; count-- { //for循环的变体形式，包含初始化语句、比较条件语句记忆对count变量执行递减运算的后置语句。
		fmt.Println(count)
	}
	fmt.Println(count) //count变量处于作用域内
}
