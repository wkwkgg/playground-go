package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("first: ", i)       // => 10
	defer fmt.Println("defer: ", i) // => 10
	i += 1
	fmt.Println("last : ", i) // => 11
}
