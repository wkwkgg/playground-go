package main

import "fmt"

func main() {
	var i int     // => 0
	var f float64 // => 0
	var b bool    // => false
	var s string  // => ""

	fmt.Printf("%v %v %v %q", i, f, b, s)
}
