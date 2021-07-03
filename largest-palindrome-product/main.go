package main

import (
	"fmt"
	"math"
)

func isPalindrome(n interface{}) bool {
	s := fmt.Sprintf("%v", n)
	r := []byte{}

	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}

	return s == string(r)
}

func LargestPalindromeByDigit(d int) int {
	largest := 0

	start := int(math.Pow10(d - 1))
	end := int(math.Pow10(d) - 1)

	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			p := i * j
			if isPalindrome(p) && p > largest {
				largest = p
			}
		}
	}

	return largest
}

func main() {

	fmt.Println(LargestPalindromeByDigit(2))
	fmt.Println(LargestPalindromeByDigit(3))
}
