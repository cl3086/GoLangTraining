package main

import "fmt"

func main() {
	// we are declaring, but not assigning. They will be assigned to their zero-value
	// int: 0, string: "", float64: 0, bool: false
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}
