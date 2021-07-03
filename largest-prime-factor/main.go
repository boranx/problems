package main

import (
	"fmt"
)

func trivialDivision(n int) []int {
	var prime []int
	f := 2 // factor

	for n > 1 {
		if n%f == 0 {
			prime = append(prime, f)
			n /= f
		} else {
			f += 1
		}
	}

	return prime
}

func max(l []int) int {
	max := 0
	for i := 0; i < len(l); i++ {
		if l[i] > max {
			max = l[i]
		}
	}

	return max
}

func main() {
	//calculate the largest for 600851475143
	fmt.Println(max(trivialDivision(600851475143)))
}
