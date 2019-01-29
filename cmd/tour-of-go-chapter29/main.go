package main

import "fmt"

// 슬라이스는 재분할 할 수도 있고, 같은 배열을 가리키는(point) 새로운 슬라이스를 만들 수 도 있습니다.

func main() {
	p := []int{2, 3, 5, 7, 11, 13}

	// 배열 전체 출력
	fmt.Println("p ==", p)

	//1부터 4까지
	fmt.Println("p[1:4] ==", p[1:4])

	// 0부터 3까지
	fmt.Println("p[:3] ==", p[:3])

	// 4부터 끝까지
	fmt.Println("p[4:] ==", p[4:])
}
