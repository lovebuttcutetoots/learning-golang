package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	City       string
	Occupation string
}

// Method attached to Person
func (p Person) Print() {
	fmt.Printf("%s is %d years old, lives in %s, and works as a %s.\n",
		p.Name, p.Age, p.City, p.Occupation)
}

func (p Person) IsAdult() bool {
	if p.Age >= 18 {
		return true
	}

	return false
}

func (p Person) Greet() string {
	return ("Hello, my name is Slim Shady. \n")
}

func main() {

	brad := Person{
		Name:       "Brad",
		Age:        42,
		City:       "Las Vegas",
		Occupation: "loser",
	}

	alice := Person{
		Name:       "Alice",
		Age:        30,
		City:       "Seattle",
		Occupation: "winner",
	}

	people := []Person{brad, alice}

	for _, person := range people {
		person.Print()
		if person.IsAdult() {
			fmt.Printf(person.Greet())
		}
	}
}
