package main

import (
	"fmt"
	"math"
)

// Pythagorean triplet is a set of three natural numbers, a < b < c, for which
// a^2 + b^2 = c^2
// such 3,4,5 - 7,24,25 - 5,12,13

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.”
// Find the product abc.

func main() {
	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			sub := i*i + j*j
			c := float64(1000 - i - j)
			s := math.Sqrt(float64(sub))
			if s == c {
				fmt.Println(i, j, c) // 200 375 425
				return
			}
		}
	}
}
