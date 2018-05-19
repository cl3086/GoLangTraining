package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a - ", a)
	// b is not being used - invalid code
}

// variables that are unused will throw an error
// a blank identifier tells Go that a variable will not be used
