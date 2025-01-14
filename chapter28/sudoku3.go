package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

type Gride [rows][columns]int8

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) //将错误转化为字符串
	}
	return strings.Join(s, ",")
}

var ( //Go语言会使用带有有Err前缀的变量存储错误信息。
	ErrBounds = errors.New("out of boiunds")
	ErrDigit  = errors.New("invalid digit")
)

func (g *Gride) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
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

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func main() {
	var g Gride
	err := g.Set(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}
