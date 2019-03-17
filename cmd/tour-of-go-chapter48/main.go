package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

/* 인터페이스는 메소드의 집합입니다.
Abser는 Abs() 를 구현한 타입이면 뭐든 적용가능합니다.
*/

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // 메소드 구현됨
	a = &v // 메소드 구현됨
	// a = v  // 구현된 메소드 없음
	// implement Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
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
