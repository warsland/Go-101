package main

import "fmt"

func main() {
	var room = "lake"
	switch { //将比较表达式放置在单独的分支内
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough //下降至下一个分支，跳过下一个分支的比较表达式。
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
	/* 将room设置为每一个分支的比较对象。
		switch { //将比较表达式放置在单独的分支内
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough //下降至下一个分支，跳过下一个分支的比较表达式。
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
	*/

}
