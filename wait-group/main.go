package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			fmt.Print(i)
		}(i)
	}

	waitGroup.Wait() // different orders for each execution like 9012345678, 9731248506 because there's race between routines
}
