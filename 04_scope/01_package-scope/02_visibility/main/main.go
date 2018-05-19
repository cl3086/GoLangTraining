package main

import (
	"fmt"

	"github.com/cl3086/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

// visibility is more the idea of whether something is exported or not from
// a package
