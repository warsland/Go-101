package main

import "fmt"

func main() {
	const lightSpeed = 299792
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That's ", days, "days of travel at light speed")
}
