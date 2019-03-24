package main

import (
	"fmt"
	"time"
)

/* 고루틴은 go 런타임에 관리 되는 경량 쓰레드 입니다.
고루틴에서 f, x, y, z가 평가되고 새로운 고루틴에서 f 가 수행됩니다.
고루틴은 동일한 주소 공간에서 실행되므로 공유되는 자원으로 접근은 반드시 동기화 되어야 합니다
golang의 sync 패키지가 동기화를 위한 기능을 제공합니다.
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
