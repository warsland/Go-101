package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello,world", "a bad apple", "goodbye all"} {
		{
			downstream <- v
		}
	}
	downstream <- ""
}
func filterGopher(upstream, downstream chan string) {
	item := <-upstream
	if item == "" {
		downstream <- ""
		return
	}
	if !strings.Contains(item, "bad") {
		downstream <- item
	}

}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
