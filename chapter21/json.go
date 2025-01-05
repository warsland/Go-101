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
		Lat, Long float64
	}
	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
