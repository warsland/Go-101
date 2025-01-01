package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terestirial := planets[0:4:4] //长度为4，容量为4
	fmt.Println(terestirial)
	terestirial = append(terestirial, "Ceres")
	fmt.Println(terestirial)
	fmt.Println(planets)
	terestirial = planets[0:4] //长度为4，容量为8
	terestirial = append(terestirial, "Ceres")
	fmt.Println(terestirial)
	fmt.Println(planets)
}
