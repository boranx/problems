package main

import (
	"fmt"
)

func divideable(s, start, end int) bool {
	for start <= end {
		if !(s%start == 0) {
			return false
		}

		start++
	}

	return true
}

func iterate(max, s, e int) int {
	found := 0
	for i := 1; i <= max; i++ {
		if divideable(i, s, e) && i > found {
			found = i
		}
	}

	return found
}

func main() {
	fmt.Println(iterate(2510, 1, 10)) // 0
	fmt.Println(iterate(3000, 1, 10)) // 2520

	fmt.Println(iterate(300000000, 1, 20)) // 232792560

}
