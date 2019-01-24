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
	fmt.Println(Vertex{1, 2})
}
