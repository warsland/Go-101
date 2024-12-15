package main

import "fmt"

func main() {
	fmt.Println("Ther is a cavern enter here an a pth to the east.")
	var command = "go inside"
	switch command { //将命令与多个分支进行比较。（单个值和单个值比较）
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside": //使用逗号可以分隔多个可选值。（单个值和多个值比较）
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get there.")
	}
}
