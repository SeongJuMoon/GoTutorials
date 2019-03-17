package main

import (
	"fmt"
	"time"
)

/* run() 리턴 타입이 error 일때
error type안에 Error 메소드를 구현함으로서 나만의
에러 처리를 구현할 수 있습니다.

*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
