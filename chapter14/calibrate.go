package main

import "fmt"

type kelvin float64
type sensor func() kelvin //sensor函数类型

func realSensor() kelvin {
	return 0
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { //声明并返回匿名函数
		return s() + offset
	}
}
func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor()) //调用函数并答应
}
