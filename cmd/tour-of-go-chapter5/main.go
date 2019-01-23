package main

import "fmt"

// /*
// 	GoLang 에선 타입 선언을 변수명 뒤에 합니다.
// */

func add(x int, y int) int {
	return x + y
}

func main() {

	fmt.Println(add(42, 13))

}
