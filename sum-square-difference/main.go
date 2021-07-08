package main

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.”

import "fmt"

func main() {
	sumSquare := 0
	squareOfSum := 0

	for i := 1; i <= 100; i++ {
		sumSquare += i * i
		squareOfSum += i
	}
	squareOfSum *= squareOfSum

	fmt.Println(sumSquare)
	fmt.Println(squareOfSum)

	// Print the diff
	fmt.Printf("diff: %d", squareOfSum-sumSquare)

}
