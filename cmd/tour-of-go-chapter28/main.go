package main

import "fmt"

/*
 []Type은 타입 T를 가지는 요소들의 집합입니다.
  슬라이스는 배열의 index(point)를 가르킴과 동시에 배열의 길이를 가지고 있습니다.
*/

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}
