package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 151)
}

func realSensor() kelvin {
	return 0
}

func main() {
	var sensor func() kelvin //声明一个不接受任何形参并值返回一个kelvin值的函数
	sensor = fakeSensor      //变量sensor的值是函数本身，而非函数结果
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
}
