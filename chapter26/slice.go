package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	plantes := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&plantes)
	fmt.Println(plantes)
}
