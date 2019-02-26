package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/* 연습 문제에서 제공해주는 패키지가 아래로 변경되었다

 */

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, word := range strings.Fields(s) {
		ret[word] += 1
	}

	return ret

}

func main() {
	wc.Test(WordCount)
	fmt.Println()
}
