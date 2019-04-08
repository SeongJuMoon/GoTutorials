package calc

// 가변 인자를 받아 반복의 합을 반환
func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}
