package main

import "fmt"

/*
 Go에 함수 역시 클로저로 동작한다.

 adder 함수의 경우 adder()의 내부함수를 리턴한다.
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i))
	}
}
