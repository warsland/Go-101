package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

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
	terraform(plantes)
	fmt.Println(plantes) //terraform函数修改的是planets数组的副本，不会影响planets数组的值
}
