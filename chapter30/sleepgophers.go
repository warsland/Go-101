package main

import (
	"fmt"
	"time"
)

func sleepGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepGopher() //启动goroutine
	}

	time.Sleep(4 * time.Second)
} //所有goroutine将在程序运行时结束。
