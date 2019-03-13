package main

import (
	"fmt"
	"math"
)

/*
 메소드는 구조체 뿐만이 아닌 타입에도 붙일 수 있습니다
 다음 예제에서는 float64를 MyFloat로 정의하여 myfloat의 대해
 Abs를 적용한 예제 입니다.
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
