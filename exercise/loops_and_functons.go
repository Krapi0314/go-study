package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, zz := 1.0, 0.0
	
	for math.Abs(z-zz) > 1e-8 {
		z, zz = z - (z*z - x) / (2*z), z
	}
	
	return z
}

func main() {
	n := 2.
	fmt.Println(Sqrt(n))
	fmt.Println(Sqrt(n) == math.Sqrt(n))
}
