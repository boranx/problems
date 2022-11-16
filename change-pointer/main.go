package main

import "fmt"

func main() {
	v := 5
	p := &v
	fmt.Println(*p) // 5

	changePointer(p)
	fmt.Println(*p) // 3
}

func changePointer(p *int) {
	v := 3
	*p = v
}
