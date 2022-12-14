package main

import (
	"sort"
	"strconv"
)

func find_elements_with_odd_number_of_occurrences(s []int) []int {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	updated_list := []int{}
	for i := 0; i < len(s); i++ {
		if i+2 > len(s) {
			updated_list = append(updated_list, s[i])
			continue
		}
		if s[i] == s[i+1] {
			// fmt.Println("match!", s[i], s[i+1])
			i++ // bump the counter two times because 2 items count as one peer
		} else {
			updated_list = append(updated_list, s[i])
		}
	}

	return updated_list
}

func createIntSlice(nums []string) []int {
	var r []int
	for _, v := range nums {
		i, _ := strconv.Atoi(v)
		r = append(r, i)
	}
	return r
}
