package main

import "fmt"

// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

// Solution:
// Zeller's congruence algorithm:
func dayOfWeek(year, month, day int) int {
	m := (month - 3 + 4800) % 4800
	y := (year + m/12) % 400
	m %= 12
	return (y + y/4 - y/100 + (13*m+2)/5 + day + 2) % 7
}

func main() {
	count := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if dayOfWeek(y, m, 1) == 0 { // Sunday
				count++
				// fmt.Println(1, "/", m, "/", y)
			}
		}
	}

	fmt.Println(count)
}
