package main

import "fmt"

func main() {
	c := 'a'
	c += 3
	if c > 'z' {
		c -= 26
	}
	fmt.Printf("%c\n", c)
}
