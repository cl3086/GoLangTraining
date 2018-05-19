package main

import "fmt"

// package level, so it is accessible in the block levels below
var x = 0

func increment() int {
	// x++ increments x
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
