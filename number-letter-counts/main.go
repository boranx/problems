package main

// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

import (
	"fmt"
	"strings"
)

const single string = "one,two,three,four,five,six,seven,eight,nine"
const special string = "ten,eleven,twelve,thirteen,fourteen,fifteen,sixteen,seventeen,eighteen,nineteen"
const doubles string = "twenty,thirty,forty,fifty,sixty,seventy,eighty,ninety"
const triples string = "hundred"

func to100() []string {
	singles := strings.Split(single, ",")
	specials := strings.Split(special, ",")
	doubles := strings.Split(doubles, ",")
	total := append(singles, specials...)

	for i := 0; i < len(doubles); i++ {
		for j := 0; j < len(singles); j++ {
			if j == 0 {
				total = append(total, doubles[i])
			}
			generated := fmt.Sprintf("%s%s", doubles[i], singles[j])
			total = append(total, generated)
		}
	}

	return total
}

func main() {
	singles := strings.Split(single, ",")

	total := to100()
	chunk := to100()
	for i := 0; i < len(singles); i++ {
		hundred := fmt.Sprintf("%s%s", singles[i], triples)
		for j := 0; j < len(chunk); j++ {
			if j == 0 {
				total = append(total, hundred)
			}
			generated := fmt.Sprintf("%sand%s", hundred, chunk[j])
			total = append(total, generated)
		}
	}

	// fmt.Println(total)
	// count how many letter was used
	counter := 0
	for i := 0; i < len(total); i++ {
		for range total[i] {
			counter++
		}
	}

	// don't forget "onethousand" ;)
	fmt.Println(counter + 11) // 21124
}
