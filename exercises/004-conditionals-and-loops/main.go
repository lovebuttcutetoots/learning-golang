package main

import "fmt"

func checkAge(name string, age int) string {
	// note: go style prefers not using an else after return
	if age < 13 {
		return name + " is a child."
	}

	if age <= 17 {
		return name + " is a teenager."
	}

	return name + " is an adult."
}

func main() {
	name := "Brad"
	age := 42

	fmt.Println(checkAge(name, age))

	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}

	names := []string{"Brad", "Alice", "Sam"}

	for _, name := range names {
		fmt.Printf("Hello, %s\n", name)
	}
}
