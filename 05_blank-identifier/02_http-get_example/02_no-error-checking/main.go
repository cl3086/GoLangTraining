package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

// the _ is an example of a blank identifier, an error is returned there, but
// we're saying that we don't quite care about it/not going to use it
