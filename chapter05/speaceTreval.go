package main

import (
	"fmt"
	"math/rand"
)

var enterpriseName = []string{"Speace Adventures", "SpeaceX          ", "Virgin Galactic  "}

const distance = 62100000

func main() {
	fmt.Println("太空航行公司        飞行天数  飞行类型  价格（百万美元）")
	for count := 0; count < 10; count++ {
		speed := rand.Intn(15) + 16
		returned := rand.Intn(2)
		time := distance / (speed * 24 * 60 * 60)
		cost := speed + 20
		re := []string{"单程", "往返"}
		if returned == 1 {
			cost *= 2
		}
		fmt.Printf("%v%8v%9v%12v\n", enterpriseName[rand.Intn(3)], time, re[returned], cost)
	}
}
