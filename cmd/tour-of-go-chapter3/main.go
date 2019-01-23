package main

import (
	"fmt"
	"math"
)

// math 패키지에서 함수를 불러올 수 있습니다.

func main() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}
