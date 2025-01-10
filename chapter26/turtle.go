package main

import "fmt"

type location struct {
	x, y int
}

type turtle struct {
	location
}

func (t *turtle) toLeft() {
	t.location.x = -1
}

func (t *turtle) toRight() {
	t.location.x = +1
}
func (t *turtle) toUp() {
	t.location.y = +1
}
func (t *turtle) toDown() {
	t.location.y = -1
}

func (t turtle) show() {
	fmt.Printf("Now tutle at %v, %v\n", t.x, t.y)
}

func main() {
	turtle := turtle{
		location{x: 0, y: 0},
	}
	turtle.show()
	turtle.toLeft()
	turtle.show()
	turtle.toUp()
	turtle.show()
	turtle.toDown()
	turtle.show()
	turtle.toRight()
	turtle.show()
}
