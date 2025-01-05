package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	type location struct {
		Lat  float64 `json:"latitude"` //使用结构标签改变输出
		Long float64 `json:"longitude"`
	}
	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
