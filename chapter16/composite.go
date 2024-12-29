package main

import "fmt"

func main() {
	plantes := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uramis",
		"Neptune",
	}
	for i := 0; i < len(plantes); i++ {
		fmt.Println(plantes[i])
	}

}
