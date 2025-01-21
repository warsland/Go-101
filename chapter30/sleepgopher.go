package main

import (
	"fmt"
	"time"
)

func sleepGopher() {
	time.Sleep(4 * time.Second)
	fmt.Println("...snore...")
}

func main() {
	go sleepGopher() //启动goroutine
	time.Sleep(4 * time.Second)
} //所有goroutine将在程序运行时结束。
