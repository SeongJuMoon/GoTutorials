package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func walker(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {

	t1TreeChanel := make(chan int)
	t2TreeChanel := make(chan int)

	go walker(t1, t1TreeChanel)
	go walker(t2, t2TreeChanel)

	for x := range t1TreeChanel {
		if v := <-t2TreeChanel; x != v {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
