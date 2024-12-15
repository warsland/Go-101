package main

import "fmt"

func main() {
	var command = "go east"
	if command == "go east" { //检查变量conmmand的值是否为go east.
		fmt.Println("You haed further up the mountain.")
	} else if command == "go inside" { //第一次检查为假后，检查变量command的值是否为go inside。
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else { //如果前两个检查为假，则执行第三个分支。
		fmt.Println("Didn't quire get there.")
	}
}
