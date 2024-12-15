package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10
	for count > 0 { //设置循环条件
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- //每次循环之后将计数减1，避免出现无限循环。
	}
	fmt.Println("Liftoff!")
}
