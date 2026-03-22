package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Value receiver (does NOT modify original)
func (p Person) BirthdayValue() {
	p.Age++
}

// Pointer receiver (modifies original)
func (p *Person) BirthdayPointer() {
	p.Age++
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func main() {

	person := Person{
		Name: "Brad",
		Age:  42,
	}

	fmt.Printf("Original age: %d\n", person.Age)

	person.BirthdayValue()
	fmt.Printf("After BirthdayValue(): %d\n", person.Age)

	person.BirthdayPointer()
	fmt.Printf("After BirthdayPointer(): %d\n", person.Age)

	person.SetAge(50)
	fmt.Printf("Updated age: %d.\n", person.Age)
}
