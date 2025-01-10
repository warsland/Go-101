package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type laser int

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (l *laser) talk() string {
	return strings.Repeat("pew", int(*l))
}

func main() {
	pew := laser(2)
	shout(pew)
}
