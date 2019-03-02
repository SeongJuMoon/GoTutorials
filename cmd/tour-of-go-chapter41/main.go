package main

import (
	"fmt"
	"runtime"
)

// switch case의 코드 실행이 끝나면 저절로 break한다

func main() {

	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(" OS X.")
	case "linux":
		fmt.Println(" Linux..")
	default:
		fmt.Printf(" %s.", os)
	}

}
