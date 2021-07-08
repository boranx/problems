package main

// “2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?”

import (
	"fmt"
)

// Warning!: Brute Force

func divideable(d, e int) bool {
	i := 1
	for i <= e {
		if !(d%i == 0) {
			return false
		}

		i++
	}

	return true
}

func iterate(max int) int {
	found := 1
	for {
		if divideable(found, max) {
			return found
		}

		found++
	}
}

func main() {
	fmt.Println(iterate(10)) // 2520
	fmt.Println(iterate(20)) // 232792560

}
