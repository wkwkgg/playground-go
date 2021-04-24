package main

import (
	"fmt"
	"strconv"
)

type List []interface{}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + "- age:" + strconv.Itoa(p.age) + ")"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hi"
	list[2] = Person{"Dennis", 79}

	for idx, elem := range list {
		switch val := elem.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", idx, val)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", idx, val)
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s\n", idx, val)
		default:
			fmt.Printf("list[%d] is of a different type\n", idx)
		}
	}
}
