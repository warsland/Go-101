package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type Gride [rows][columns]int8

func (g *Gride) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of boiunds")
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func main() {
	var g Gride
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occureed: %v.\n", err)
		os.Exit(1)
	}
}
