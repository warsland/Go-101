package main

import (
	"fmt"
	"strings"
)

var t interface { //声明接口
	talk() string
}

type martian struct{}

func (m martian) talk() string { //满足接口的要求
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

func main() {
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
}
