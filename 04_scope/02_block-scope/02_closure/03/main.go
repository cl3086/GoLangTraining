package main

import "fmt"

func main() {
	x := 0

	// increment is assigned an anonymous function (no name), it is now increment
	// to create a function inside a function, it must be an anonymous function
	// x is in the outer scope, it is accessible inside increment()
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
