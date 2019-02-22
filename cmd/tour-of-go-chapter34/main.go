package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// map 은 키와 밸류 타입을 지정해서 key값을 넣을 수 있다
// 맵의 경우 new를 사용하지 않고 make function으로

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
