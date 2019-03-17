package main

/*
OS.stdin에서 내부적으로 정의된 writer가 암시적으로 구현되어있어서
writer 인터페이스를 적용한다
*/

import (
	"fmt"
	"os"
)

type Reader interface {
	Reader(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	w = os.Stdout

	fmt.Fprintf(w, "Hello, writer\n")
}
