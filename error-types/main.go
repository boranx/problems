package main

import (
	"fmt"
)

type errorString struct{}

func (e *errorString) Error() string { return "" }

func checkErr(err error) {
	// interface (type|value)
	// fmt.Println(type == nil)
	fmt.Println(err == nil)
}

func main() {

	var e1 error
	checkErr(e1) // true

	var e *errorString
	checkErr(e) // false

	e = &errorString{}
	checkErr(e) // false

	e = nil
	checkErr(e) // false
}
