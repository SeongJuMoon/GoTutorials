package main

import (
	"fmt"
	"math/cmplx"
)

/*
bool
string

int  int8  int16  int32  int64  (+ - 정수)
uint uint8 uint16 uint32 uint64 uintptr (+ 정수)

byte // uint8의 다른 이름(alias)

rune // int32의 다른 이름(alias)
     // 유니코드 코드 포인트 값을 표현합니다.

float32 float64 (실수)

complex64 complex128 (복소수)
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n" // 서식 문자 타입과 값 출력
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
