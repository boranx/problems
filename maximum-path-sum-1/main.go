package main

import (
	"fmt"
	"strconv"
	"strings"
)

// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

// 3
// 7 4
// 2 4 6
// 8 5 9 3

// That is, 3 + 7 + 4 + 9 = 23.

// Find the maximum total from top to bottom of the triangle below:

// 75
// 95 64
// 17 47 82
// 18 35 87 10
// 20 04 82 47 65
// 19 01 23 75 03 34
// 88 02 77 73 07 63 67
// 99 65 04 28 06 16 70 92
// 41 41 26 56 83 40 80 70 33
// 41 48 72 33 47 32 37 16 94 29
// 53 71 44 65 25 43 91 52 97 51 14
// 70 11 33 28 77 73 17 78 39 68 17 57
// 91 71 52 38 17 14 91 43 58 50 27 29 48
// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

// Solution: Using Dynamic Programming, starting from bottom to up
// Calculate best patterns

var triangle = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func main() {
	var triangle_leaves [][]int
	structure := strings.Split(triangle, "\n")
	for _, v := range structure {
		temp := strings.Fields(v)
		triangle_leaves = append(triangle_leaves, sliceAtoi(temp))
	}

	for i := len(triangle_leaves) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle_leaves[i][j] += max(triangle_leaves[i+1][j], triangle_leaves[i+1][j+1])
		}
	}

	// Debug: print out whole sum by each level if you wanna see how it narrowed
	// fmt.Println(triangle_leaves)
	fmt.Println(triangle_leaves[0][0]) // 1074
}

func sliceAtoi(s []string) []int {
	l := make([]int, 0, len(s))
	for _, a := range s {
		i, err := strconv.Atoi(a)
		if err != nil {
			return nil
		}
		l = append(l, i)
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
