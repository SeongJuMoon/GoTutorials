package main

import (
	"fmt"
	"math/cmplx"
)

/*
 complex64 타입과 complex128 타입을 통해서 Golang의 복소수 지원 기능을 알아봅니다

 아래의 예는 뉴턴 메서드를 적용하여 다음 반복을 수행합니다
*/

func Cbrt(x complex128) complex128 {

	if x == 0 {
		return 0
	}

	z := x

	// 루핑되어 계속 쪼개지는 z^3가 x의 값과 같아질 때 까지 반복한다.
	for cmplx.Pow(z, 3) != x {
		z = z - (z*z*z-x)/(3*z*z)
	}

	return z

}

func main() {
	fmt.Println(Cbrt(2), "=?", cmplx.Pow(2, 1./3))
}
