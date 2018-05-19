package main

import "fmt"

func wrapper() func() int {
	// wrapper is a function that returns a function (that returns an int)
	// go is a first order language
	// so anytime this function is called, it accesses x (closure)
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	// note, wrapper returns a new function every time (so x is initially 0)
	increment1 := wrapper()
	fmt.Println(increment1())
	fmt.Println(increment1())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
