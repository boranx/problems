package main

// The following iterative sequence is defined for the set of positive integers:

// n → n/2 (n is even)
// n → 3n + 1 (n is odd)

// Using the rule above and starting with 13, we generate the following sequence:

// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

// Which starting number, under one million, produces the longest chain?

// NOTE: Once the chain starts the terms are allowed to go above one million.

func sequence(n int) []int {
	var l []int
	for {
		if n == 1 {
			l = append(l, 1)
			return l
		}
		if n%2 == 0 {
			l = append(l, n)
			n = n / 2
		} else {
			l = append(l, n)
			n = 3*n + 1
		}
	}
}

func main() {
	const UPPER_LIMIT = 1_000_000
	n, max := 0, 0

	for i := 1; i < UPPER_LIMIT; i++ {
		if len(sequence(i)) > max {
			max = len(sequence(i))
			n = i
		}
	}

	println(n, " -> len():", max) // 837799  -> len(): 525
}
