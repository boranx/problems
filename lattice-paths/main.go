package main

import (
	"math/big"
)

// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
// https://projecteuler.net/problem=15
// How many such routes are there through a 20×20 grid?

// f(h, w) = f(h, w - 1) + f(w, h - 1) // didn't work because it's time consuming
func paths(i, j int) int {
	if i == 1 || j == 1 {
		return 1
	}

	return paths(i-1, j) + paths(i, j-1)
}

// (m+n m)=(m+n n)= (m+n)! / m!*n!
func path_v2(n int) string {
	b := factorial(n)
	b.Mul(b, b)

	return new(big.Int).Div(factorial(2*n), b).String()
}

func main() {
	println(paths(2, 2)) // 2
	println(paths(3, 3)) // 6
	// println(paths(20, 20)) // ∞
	println(path_v2(20))
}

func factorial(n int) *big.Int {
	f := big.NewInt(1)
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}
