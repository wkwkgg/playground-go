package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
