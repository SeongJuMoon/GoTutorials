package main

import (
	"fmt"
	"math"
)

/* 고에는 class가 존재하지 않습니다.
   하지만 메소드를 구조체에 붙일 수 있습니다.
   메소드 수신자는 메소드명 이전에 인자로 들어갑니다.
*/
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
