package main

import (
	"fmt"
	"math"
)

func main() {
	// 함수로 아래와 같이 인자로서 사용가능

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
