package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury" //数组中的元素赋值
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] //访问数组中的元素
	fmt.Println(earth)

	fmt.Println(len(planets)) //获得数组的长度
	fmt.Println(planets[3] == "")

}
