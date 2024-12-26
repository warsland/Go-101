package main

import "fmt"

type celsius float64
type fahrenherit float64
type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenherit() fahrenherit {
	return fahrenherit(k - 459.67)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenherit() fahrenherit {
	return fahrenherit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenherit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenherit) kelvin() kelvin {
	return kelvin(f + 459.67)
}

func main() {
	var k kelvin = 0.0
	var c celsius = -40.0
	var f fahrenherit = 100.0
	fmt.Println(k)
	fmt.Println(k.celsius())
	fmt.Println(k.fahrenherit())
	fmt.Println(c)
	fmt.Println(c.kelvin())
	fmt.Println(c.fahrenherit())
	fmt.Println(f)
	fmt.Println(f.celsius())
	fmt.Println(f.kelvin())
}
