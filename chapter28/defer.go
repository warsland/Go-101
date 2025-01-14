package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name) //创建文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "Errors are values.")
	if err != nil { //写入文件失败
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle tme gracefully.")
	return err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
