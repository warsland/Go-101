package main

import "fmt"

func main() {
	num := make([]int, 0, 5)
	capa := 5
	for i := 0; i < 200; i++ {
		num = append(num, i+1)
		newCapa := cap(num)
		if capa != newCapa {
			fmt.Printf("old capacity: %v, new capacity: %v, now capacity: %v\n", capa, newCapa, len(num))
			capa = newCapa
		}

	}
}
