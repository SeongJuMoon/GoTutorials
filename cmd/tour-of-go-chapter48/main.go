package main

import (
	"fmt"
	"math"
)

/*
 인터페이스란 메소드 집합이므로
 그 메소드를 구현한 타입은 모두 인터페이스 값이 될 수 있습니다.
*/

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser

	f := MyFloat64(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // 1.4141414 implements abs()
	a = &v // 5 implements abs()
	//	a = v // not implements abs

	fmt.Println(a.Abs())
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
