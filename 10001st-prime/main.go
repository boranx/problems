package main

//“By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?”

import "fmt"

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func findTo(n int) []int {
	var primes []int

	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}

	}

	return primes
}

func last(l []int) int {
	return l[len(l)-1]

}

func main() {
	fmt.Println(last(findTo(6)))
	fmt.Println(last(findTo(10_001)))
}
