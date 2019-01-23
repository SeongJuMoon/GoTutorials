package main

import (
	"fmt"
	"math"
)

/*
	프로그램은 main 패키지에서부터 실행을 시작합니다.
	이 프로그램은 `"fmt"`와 "math" 패키지를 import 해서 사용하고 있습니다.
	패키지 이름은 디렉토리 경로의 마지막 이름을 사용하는 것이 규칙입니다.
*/

func main() {
	fmt.Println("happy", math.Pi, "Day")
}
