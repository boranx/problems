package main

import "math"

// The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

// Let us list the factors of the first seven triangle numbers:

//  1: 1
//  3: 1,3
//  6: 1,2,3,6
// 10: 1,2,5,10
// 15: 1,3,5,15
// 21: 1,3,7,21
// 28: 1,2,4,7,14,28

// We can see that 28 is the first triangle number to have over five divisors.
// What is the value of the first triangle number to have over five hundred divisors?

func main() {
	i := 1
	for {
		triangle := triangle(i)
		divisor := numOfDivisior(triangle)
		if divisor > 500 {
			println(triangle, divisor) // 76576500 576
			return
		}

		i++
	}
}

func triangle(n int) int {
	return n * (n + 1) / 2
}

// better than trivial division:
// for each a,b : factors of n,
// a x b = n (For every exact divisor up to the square root, there is a corresponding divisor above the square root)
// if a < b then a < √n and b > √n or a = b = √n
func numOfDivisior(n int) int {
	numDivisors := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			numDivisors += 2
		}
	}

	return numDivisors
}
