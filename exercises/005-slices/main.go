package main

import "fmt"

func printNames(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	names := []string{"Brad", "Alice", "Sam"}

	fmt.Println("Original names:")
	printNames(names)

	names = append(names, "Jordan")

	fmt.Println("\nAfter append:")
	printNames(names)

	fmt.Printf("\nLength of Names: %d\n", len(names))
	fmt.Printf("First Name: %s\n", names[0])
	fmt.Printf("Last Name: %s\n", names[len(names)-1])
}
