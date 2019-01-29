package main

import "fmt"

type Vertex struct {
	X, Y int
}

// case 1
var t *Vertex = new(Vertex)

func main() {
	// case 2
	v := new(Vertex) // new를 사용할 경우 모든 필드가 숫자인 경우 0 레퍼런스 타입의 경우 nil이 할당된 포인터를 반환합니다.

	v.X, v.Y = 11, 9

	fmt.Println(v)

}
