package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) //必须要进行类型转换
}

func main() {
	var k kelvin = 294.0 //实参必须为kelvin类型
	c := kelvinToCelsius(k)
	fmt.Print(k, "K IS ", c, "℃\n")
}
