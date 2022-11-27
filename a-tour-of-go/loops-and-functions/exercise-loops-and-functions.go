package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const precision = 0.0001
	y, z := 0.0, 1.0
	for math.Abs(y-z) >= precision {
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	x := 200.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
