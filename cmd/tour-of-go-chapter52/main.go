package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

// h 구조체의 HTTP 서브 구현한다

// responsewrite 함수 호출 시점의 생성 이후 fPrint로 결과 값 작성
func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// 이후 http server Listen
func main() {
	var h Hello
	http.ListenAndServe("localhost:5000", h)
}
