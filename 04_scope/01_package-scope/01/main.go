package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}

// this is an example of package level scope
// the variable x is accessible to the entire package (ANY file within the directory!)
// the functions are enclosed within the package level scope

// y is initialized in a block level scope. This means y is only accessible within
// the enclosing block (func main)
// y is not accessible in foo
// with block level scope, order matters (have to initialize y before we can use it)
