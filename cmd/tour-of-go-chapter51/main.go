package main

import (
	"fmt"
)

/*
 */

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %f", e)
}

// func Sqrt(f float64) (float64, error) {
// 	return 0, nil
// }

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := 1.0

	for v := 0; v < 5; v++ {
		z = z - (z*z-f)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
