package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

//type sensor func() kelvin //使用type关键字声明函数类型

func measureTemperature(samples int, sensor func() kelvin) { //measureTemperature接受另一个函数作为其第二个形参
	//func measureTemperature(samples int, s sensor) //使用type关键字声明函数类型给后的代码声明
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)

	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor) //将函数明传递给另一个函数
}
