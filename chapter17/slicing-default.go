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
		"Uranus",
		"Neptune",
	}

	terrestrial := plantes[:4] //创建切片
	gasGiants := plantes[4:6]
	iceGiants := plantes[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	allPlantes := plantes[:]
	fmt.Println(allPlantes)
}
