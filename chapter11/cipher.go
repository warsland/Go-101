package main

import (
	"fmt"
	"strings"
)

var plainText = "your message goes here"
var keyword = "GOLANG"

func main() {
	var encryption string

	newText := strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	fmt.Println("未加密的文本为：", newText)
	for i := 0; i < len(newText); i++ {
		num := len(keyword)
		shitf := keyword[i%num]
		c := (newText[i]-'A'+shitf)%26 + 'A'
		encryption += string(c)
	}
	fmt.Println("加密后的文本为：", encryption)
}
