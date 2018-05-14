package stringutil

func Reverse(s string) string {
	// reverseTwo is defined in the reverseTwo.go file
	// it is accessible within the package
	return reverseTwo(s)
}

/*
  go build reverse.go reverseTwo.go will not produce an output file
  go install will place the package insdie the pkg directory of the workspace
*/
