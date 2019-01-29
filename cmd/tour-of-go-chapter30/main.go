package main

import "fmt"

// 슬라이스는 make 함수로 만들 수 있습니다. 이렇게 생성된 슬라이스는 0을 할당한 배열을 생성하고, 그것을 참조(refer)합니다.

// cap() 배열(슬라이스)의 용량을 알아 내는 것
// len() 배열의 길이를 알아 내는 것
// make(타입, 채울 값, 용량)

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d,cap=%d%v\n", s, len(x), cap(x), x)
}
