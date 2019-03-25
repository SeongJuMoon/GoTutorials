package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 채널에 sum 변수를 전송
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // 채널 선언
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 채널에서 값을 가져옴

	fmt.Println(x, y, x+y)

}
