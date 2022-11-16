package main

import (
	"fmt"
	"sort"
)

// calculate pairs
// example: []int32{1, 2, 1, 2, 2, 3, 4, 3, 3, 3}
// result: 4
func countPairs(s []int32) int32 {

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	pair := 0
	fmt.Println(s) // [1 1 2 2 2 3 3 3 3 4]

	for i := 0; i < len(s); i++ {
		if i+2 > len(s) {
			break
		}
		if s[i] == s[i+1] {
			fmt.Println("match!", s[i], s[i+1])
			pair++
			i++ // bump the counter two times because 2 items count as one peer
		}

	}

	return int32(pair)
}

func main() {
	fmt.Println(countPairs([]int32{1, 2, 1, 2, 2, 3, 4, 3, 3, 3}))                         // 4
	fmt.Println(countPairs([]int32{1, 2, 1, 2, 2, 3, 4, 3, 3, 3, 3, 3, 4, 5, 6, 7, 5}))    // 7
	fmt.Println(countPairs([]int32{1, 2, 1, 2, 2, 3, 4, 3, 3, 3, 3, 3, 4, 5, 6, 7, 5, 6})) // 8
	fmt.Println(countPairs([]int32{1, 2, 3, 4, 5, 4, 7, 8, 8}))                            // 2
}
