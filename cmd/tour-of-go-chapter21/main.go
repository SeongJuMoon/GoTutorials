package main

import (
	"fmt"
	"math"
)

/*
	미적분학에서 다루는 뉴턴의 방법으로 구현된 함수와
	Sqrt함수의 결과값을 비교하는 예제입니다.
*/

func Sqrt(x float64) float64 {
	z := float64(1)
	z = z - (z*z-x)/(z*2)

	return z
}

func main() {

	for x := 1; x < 10; x++ {
		fmt.Println("Newton's method ", Sqrt(float64(x)))
		fmt.Println("Math", math.Sqrt(float64(x)))
	}

}
