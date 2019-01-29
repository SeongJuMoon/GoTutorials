package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // q는 vertex의 레퍼런스 변수
	r = Vertex{X: 1}  // Y는 0으로 type의 초기 값으로 초기화 됨
	s = Vertex{}      // X = 0 and Y & 0 초기화
)

func main() {
	fmt.Println(p, q, r, s)
}
