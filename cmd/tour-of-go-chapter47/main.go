package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
 고에서 포인터 리시버를 사용하는 이유는 2가지
 큰 구조체의 값을 메소드 호출때마다 복사하는 비용이 너무 커서 이를 방지하기 위하여
 다른 이유는 메소드의 수신 포인터의 값을 바꾸기 위해서 입니다.

 아래의 스케일은 v.X와 v.Y의 값을 바꾸는 메소드입니다.
 아래의 Vertex를 에서 포인터를 제거하면 값이 바뀌지 않습니다.
 그 이유는 메소드 호출시의 그 값을 바꾸는 것이 아닌 값을 복사한 후 그 값을 바꾸기 때문입니다.
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4} // 구조체 초기화
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
