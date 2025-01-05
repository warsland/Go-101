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

type location struct {
	Name string  `json:"name"`
	Lat  float64 `json:latitude`
	Long float64 `json:longitude`
}

func main() {
	locations := []location{
		{"Bradbury Landing", -4.5895, 137.4417},
		{"Columbia Meorial Sation", -14.5684, 175.472636},
		{"Chellenger Memorial Station", -1.9462, 354.4734},
	}
	for _, location := range locations {
		bytes, err := json.MarshalIndent(location, "", "\t")
		exitOnError(err)
		fmt.Println(string(bytes))
	}
}
