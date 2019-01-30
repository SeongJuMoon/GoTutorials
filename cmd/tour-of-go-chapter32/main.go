package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// for loop에서 range를 사용하면 배열(슬라이스)나 맵을 순회할 수 있습니다.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
