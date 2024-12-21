package main

import "fmt"

func main() {
	const distance = 240000000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay

	fmt.Println("Andrimeda Galaxy is", days, "light days away.")
}
