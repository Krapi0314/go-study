package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}	

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return x, ErrNegativeSqrt(x)
	}

	z, zz := 1.0, 0.0

	for math.Abs(z-zz) > 1e-8 {
		z, zz = z-(z*z-x)/(2*z), z
	}

	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
