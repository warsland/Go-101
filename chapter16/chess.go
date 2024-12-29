package main

import "fmt"

func main() {
	var board [8][8]string //定义一个8*8的二维数组

	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Print(board)
}
