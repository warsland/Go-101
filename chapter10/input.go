package main

import "fmt"

func main() {
	var input = "false"
	var output = true
	switch input {
	case "true", "yes", "1":
		fmt.Println(output)
	case "false", "no", "0":
		output = false
		fmt.Println(output)
	}

}
