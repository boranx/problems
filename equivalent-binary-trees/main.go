package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value // process node

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go FullWalk(t1, ch1)
	ch2 := make(chan int)
	go FullWalk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 { // no more nodes in trees
			break
		}
		if ok1 != ok2 { // trees with different number of nodes
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
func main() {
	a := tree.New(2)
	b := tree.New(2)

	fmt.Println("a ->", a.String())
	fmt.Println("b ->", b.String())
	fmt.Println("Trees equivalent?", Same(a, b))

}

// walk the full tree, close output channel when done
func FullWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}
