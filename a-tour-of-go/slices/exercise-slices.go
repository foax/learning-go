package main

import (
	"golang.org/x/tour/pic"
	// "fmt"
)

func pixel_value(x, y int) uint8 {
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			z := pixel_value(x, y)
			pic[y][x] = z
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
