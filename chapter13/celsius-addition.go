package main

func main() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	var warmUp float64 = 10
	temperature += celsius(warmUp)
}
