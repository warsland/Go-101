package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
		select {
		case gopherID := <-c:
			fmt.Printf("gopher %v has finished sleeping\n", gopherID)
		case <-timeout:
			fmt.Println("my patience ran out")
		}
	}
}

func sleepGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000) * int(time.Millisecond)))
	c <- id
}
