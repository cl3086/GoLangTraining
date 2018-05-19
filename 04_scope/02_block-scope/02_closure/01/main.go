package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		// x is in the outer scope so it is accessible inside here
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}
