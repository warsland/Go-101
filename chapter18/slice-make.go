package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10) //指定切片的长度为0，容量为10
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)
}
