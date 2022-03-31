package main

import (
	"math/big"

	"strconv"
)

// n! means n × (n − 1) × ... × 3 × 2 × 1

// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

// Find the sum of the digits in the number 100!

func bigFactorial(number int) *big.Int {
	if number == 1 {
		return big.NewInt(1)
	}

	return big.NewInt(1).Mul(bigFactorial(number-1), big.NewInt(int64(number)))
}

func digitSum(vector *big.Int) int {
	sum := 0
	for _, v := range vector.String() {
		vector, _ := strconv.Atoi(string(v))
		sum += vector
	}

	return sum
}

func main() {
	println(digitSum(bigFactorial(10)))  // 27
	println(digitSum(bigFactorial(100))) // 648
}
