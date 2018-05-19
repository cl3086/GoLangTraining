package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42

// the order of initialization matters (x is used before it is initialized)
// y however, is defined at the package level, so it is accessible throughout the
// entire package (so order doesn't matter here)
