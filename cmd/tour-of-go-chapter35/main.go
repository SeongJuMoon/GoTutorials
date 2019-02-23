package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 맵 역시 리터럴을 사용할 수 있으며 타입을 선언 경우에 { } 안에 값을 넣어서 초기화할 수 있다.

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
