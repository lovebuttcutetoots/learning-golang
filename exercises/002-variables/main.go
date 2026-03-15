package main

import "fmt"

func main() {

	// Full Declaration Syntax
	/*
		var name string = "Brad"
		var age int = 42
		var height float64 = 5.10
		var learningGo bool = true
	*/

	// Short Declaration Syntax
	name := "Brad"
	age := 42
	height := 5.10
	learningGo := true

	// Println Syntax
	/*
		fmt.Println("Name:", name)
		fmt.Println("Age:", age)
		fmt.Println("Height:", height)
		fmt.Println("Learning Go:", learningGo)
	*/

	// Printf Syntax
	fmt.Printf("Hello, my name is %s.\n", name)
	fmt.Printf("I am %d years old.\n", age)
	fmt.Printf("My height is %.2f feet.\n", height)
	fmt.Printf("Learning Go: %t\n", learningGo)

}
