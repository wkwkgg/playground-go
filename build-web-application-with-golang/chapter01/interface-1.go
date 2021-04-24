package main

import "fmt"

type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human
	School string
	load   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// SayHi method for Human object
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s. You can call me on %s\n", h.name, h.phone)
}

// Sing method for Human object
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la", lyrics)
}

// Guzzle method for Human object
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle ...", beerStein)
}

// overload
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.load += amount
}

func (e *Employee) SpendMoney(amount float32) {
	e.money -= amount
}

// define interfaces
type Men interface {
	SayHi()
	Sing(lyrics string)
	// Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {
	mike := Student{Human{"Mike", "222-222-XXXX", 25}, "MIT", 0.00}
	paul := Student{Human{"Paul", "333-333-YYYY", 26}, "Harvard", 100}
	sam := Employee{Human{"Sam", "444-444-ZZZZ", 36}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", "555-555-XYZZ", 37}, "Things Ltd.", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is Tome, a Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// define Men in slice
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, v := range x {
		v.SayHi()
	}
}
