package main

import "fmt"

// define functions
// func func-name(arg-name arg-type, ...) return-type { ... }

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}