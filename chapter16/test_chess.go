package main

import "fmt"

func main() {
	var board [8][8]rune

	for i := 0; i < 8; i++ {
		if i == 0 {
			for j := 0; j < 8; j++ {
				switch j {
				case 0, 7:
					board[i][j] = 'r'
				case 1, 6:
					board[i][j] = 'n'
				case 2, 5:
					board[i][j] = 'b'
				case 3:
					board[i][j] = 'k'
				case 4:
					board[i][j] = 'q'
				}

			}
		} else if i == 1 {
			for column := range board[1] {
				board[i][column] = 'p'
			}
		} else if i == 6 {
			for column := range board[1] {
				board[i][column] = 'P'
			}
		} else if i == 7 {
			for j := 0; j < 8; j++ {
				switch j {
				case 0, 7:
					board[i][j] = 'R'
				case 1, 6:
					board[i][j] = 'N'
				case 2, 5:
					board[i][j] = 'B'
				case 3:
					board[i][j] = 'K'
				case 4:
					board[i][j] = 'Q'
				}
			}
		} else {
			continue
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}
}
