package main

import "fmt"

/*
	상수는 const 키워드와 함께 변수처럼 선언합니다
	상수는 숫자,문자열,문자 등의 타입으로 선언 가능합니다.
*/

const Pi = 3.14

func main() {
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
