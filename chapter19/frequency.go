package main

import "fmt"

func main() {
	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int) //使用make函数创建映射

	for _, t := range temperature { //迭代切片，取出其中的索引和值
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}
