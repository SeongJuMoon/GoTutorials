package main

import (
	"fmt"
	"time"
)

// 같은 변수로 if then else 대신 아래와 같은 코드를 작성하면 직관적으로 코딩할 수 있다.

func main() {

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("it's am times")
	case t < 17:
		fmt.Println("it's dinner time")
	default:
		fmt.Println("let's play game")
	}

}
