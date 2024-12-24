package main

import "fmt"

var cipherText = "COSITEUIWUIZNSROCNKFD"
var keyword = "GOLANG"

func main() {
	var answer string
	//不使用range关键字
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		num := len(keyword)
		shift := keyword[i%num]
		c = ((c-'A')+shift)%26 + 'A'
		answer += string(c)
	}
	fmt.Println(answer)

	//使用range关键字
	for i, c := range cipherText {
		num := len(keyword)
		shift := keyword[i%num]
		var d = ((byte(c)-'A')+shift)%26 + 'A'
		answer += string(d)
	}
	fmt.Println(answer)
}
