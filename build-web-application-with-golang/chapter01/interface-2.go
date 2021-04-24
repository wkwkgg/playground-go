package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name, phone string
	age         int
}

func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆" + h.phone + "❱"
}

func main() {
	Bob := Human{"bob", "012-345-XXXX", 39}
	fmt.Println("This human is : ", Bob)
}
