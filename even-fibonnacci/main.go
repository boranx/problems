package main

import "fmt"

// in a fibonacci series whose values do not exceed four million, find the sum of the even-valued terms
func main() {
	a := 0
	b := 1
	const max = 4_000_000
	sum := 0

	for b < max {
		// fmt.Println(a)
		if b%2 == 0 {
			sum = sum + b
		}
		a, b = b, b+a
	}

	fmt.Println(sum)
}
