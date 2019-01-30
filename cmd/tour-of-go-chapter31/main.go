package main

import "fmt"

// 빈 슬라이스 다음과 같이 확인할 수 있습니다.

func main() {
	var z []int
	// 초기화 되지 않은 배열(슬라이스)의 길이(len)과 용량(cap)은 0입니다.

	fmt.Println(z,len(z),cap(z))

	if z == nil {
		fmt.Println("nil")
	}

}