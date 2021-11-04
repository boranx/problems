package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// What is the sum of the digits of the number 2^1000?
// https://projecteuler.net/problem=16

func exponentTo(n int) big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(2))
	}

	return *result
}

func digits(n big.Int) []int {
	l := []int{}
	chars := n.String()
	for {
		if len(chars) == 0 {
			return l
		}
		v, _ := strconv.Atoi(chars[:1])
		l = append(l, v)
		chars = chars[1:]
	}
}

func sum(l []int) int {
	if len(l) == 0 {
		return 0
	}

	return l[0] + sum(l[1:])
}

func main() {
	fmt.Println(sum(digits(exponentTo(15))))   // 26
	fmt.Println(sum(digits(exponentTo(1000)))) // 1366
}
