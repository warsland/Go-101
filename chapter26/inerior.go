package main

import "fmt"

type stats struct {
	level            int
	endurance, heath int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.heath = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main() {
	plyer := character{name: "Matthias"}
	levelUp(&plyer.stats)
	fmt.Printf("%+v\n", plyer.stats)
}
