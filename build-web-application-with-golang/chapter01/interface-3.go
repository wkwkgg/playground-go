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
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if val, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, val)
		} else if val, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, val)
		} else if val, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an Person and its value is %s\n", index, val)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
