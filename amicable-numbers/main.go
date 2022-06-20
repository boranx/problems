package main

import "math"

// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.
// https://projecteuler.net/problem=21

func sumOfDivisors(n int) int {
	sum := 0
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}

	if n == sqrt*sqrt {
		sum = sum - sqrt
	}

	return sum + 1 // add lowest divisor as we previously omit
}

func isAmicable(n int) bool {
	s := sumOfDivisors(n)

	// eliminate `perfect` numbers :)
	if s == n {
		return false
	}

	d := sumOfDivisors(s)
	if d == n {
		return true
	}

	return false
}

func main() {
	sum := 0
	for i := 2; i <= 10000; i++ {
		if isAmicable(i) {
			print(i, ",")
			sum += i // 220,284,1184,1210,2620,2924,5020,5564,6232,6368
		}
	}
	println()
	println("sum:", sum) // 31626
}
