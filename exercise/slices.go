package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := range s {
		s[y] = make([]uint8, dx)
		for x := range s[y] {
			s[y][x] = interpretCoordinate(x,y)
		}
	}
	
	return s
}

func interpretCoordinate(x, y int) uint8 {
	n := x^y
	
	return uint8(n)
}

func main() {
	pic.Show(Pic)
}
