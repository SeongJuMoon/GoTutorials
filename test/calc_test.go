package calc_test

import (
	"calc"
	"testing"
)


// 만약 테스트에 통과하지 못하면 error 메시지 출력
func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	if s ! = 6 {
		t.Error("Not accepts test")
	}
}


