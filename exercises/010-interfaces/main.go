package main

import "fmt"

// Interface
type Speaker interface {
	Speak() string
}

// Struct 1
type Person struct {
	Name string
}

// Struct 2
type Robot struct {
	ID string
}

// Struct 3
type Dog struct {
	Name string
}

// Person implements Speaker
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name + "."
}

// Robot implements Speaker
func (r Robot) Speak() string {
	return "Beep boop. My ID is " + r.ID + "."
}

// Dog implements Speaker
func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name + "."
}

// Function that accepts the interface
func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	person := Person{Name: "Brad"}
	robot := Robot{ID: "RX-78"}
	dog := Dog{Name: "Athena"}

	saySomething(person)
	saySomething(robot)
	saySomething(dog)
}
