package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.0106
	fmt.Printf("%+v\n%+v\n", bradbury, curiosity)
}
