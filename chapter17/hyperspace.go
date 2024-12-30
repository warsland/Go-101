package main

import (
	"fmt"
	"strings"
)

func hyperspeace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	plantes := []string{"Venus", "Earth", "Mars"}
	hyperspeace(plantes)
	fmt.Println(strings.Join(plantes, ""))
}
