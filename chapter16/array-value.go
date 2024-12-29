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

	plantesMark2 := plantes
	plantes[2] = "whoops"

	fmt.Println(plantes)
	fmt.Println(plantesMark2)

}
