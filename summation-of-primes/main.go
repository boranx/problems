package main

import (
	"fmt"
	"math"
)

// “The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.”

func main() {
	fmt.Println(sum(10))        // 17
	fmt.Println(sum(2_000_000)) // 142913828922
}

func sum(d int) int {
	sum := 0
	for i := 2; i < d; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	return sum
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
