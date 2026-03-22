package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	City       string
	Occupation string
}

func printPerson(p Person) {
	fmt.Printf("%s is %d years old, lives in %s, and works as a %s.\n", p.Name, p.Age, p.City, p.Occupation)
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

	sam := Person{
		Name:       "Sam",
		Age:        27,
		City:       "Denver",
		Occupation: "winner",
	}

	jordan := Person{
		Name:       "Jordan",
		Age:        35,
		City:       "Phoenix",
		Occupation: "chicken dinner",
	}

	people := []Person{brad, alice, sam, jordan}

	for _, person := range people {
		printPerson(person)
	}
}
