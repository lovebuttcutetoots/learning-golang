package main

import "fmt"

func printAges(people map[string]int) {
	for name, age := range people {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}

func main() {

	// Create a map
	// map[keyType]valueType
	people := map[string]int{
		"Brad":  42,
		"Alice": 30,
		"Sam":   27,
	}

	// Access element of the map
	fmt.Printf("Brad is %d years old.\n\n", people["Brad"])

	printAges(people)

	// Add a value to the map
	people["Jordan"] = 35

	fmt.Println()
	fmt.Println("After adding Jordan:")
	fmt.Println()

	printAges(people)

	// Prove that Sam exists in the map
	_, ok := people["Sam"]
	if ok {
		// Key exists
		fmt.Printf("Sam exists!\n")
	} else {
		// Key does not exist
		fmt.Println("Value does not exist.")
	}

	// Delete Alice
	delete(people, "Alice")

	fmt.Println()
	fmt.Println("After delete:")
	fmt.Println()

	printAges(people)
}
