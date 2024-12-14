package main

import "fmt"

func main() {
	//fmt.Println("My weight on the surface of Mars is ")//文本
	//fmt.Println(149.0 * 0.3783)//数学表达式
	//fmt.Println("lbs,and I would be ")
	//fmt.Println(41 * 365 / 687)
	//fmt.Println("years old")
	//fmt.Println("My weight on the surface of Mars is ", 149.0*0.3783, "lbs,and I would be ", 41*365/687, "years old")
	fmt.Printf("My weight on the surface of Mars is  %v lbs\n", 149.0*0.3783)
	fmt.Printf("and I would be  %v years old\n", 41*356/687)
	fmt.Printf("%-15v $%4v\n", "SpeaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Vrigin Galactic", 100)
}
