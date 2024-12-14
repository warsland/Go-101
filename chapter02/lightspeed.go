package main

import "fmt"

func main() {
	const lightSpeed = 297792
	var distance = 56000000
	var speed = 100800
	/*var (
		distance = 56000000
		speed=100800
	)*/
	//var distance,speed=56000000,100800
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/speed, "seconds")
}
