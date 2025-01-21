package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("... %v snore...\n", id)
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
