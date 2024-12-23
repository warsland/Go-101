package main

import "fmt"

func main() {
	launch := false
	luanchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for laucnch:", luanchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for laucnch:", yesNo)
}
