package main

/*
	구조체는 필드(데이터)들의 조합입니다.
*/

import "fmt"

type Vertex struct { // 구조체의 타입 별칭을 alias
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	v.X = 4
	fmt.Println(v.X)
	// 구조체의 필드는 .(dot)으로 접근할 수 있습니다.
}
