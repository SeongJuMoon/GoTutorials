package main

import (
	"fmt"
	"image"
)

// https://golang.org/pkg/image/#Image

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA)
}
