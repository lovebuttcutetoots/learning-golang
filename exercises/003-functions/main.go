package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func subtract(a, b int) int {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func main() {
	name := "Brad"

	fmt.Println(greet(name))
	fmt.Printf("2 + 3 = %d\n", add(2, 3))
	fmt.Printf("4 * 5 = %d\n", multiply(4, 5))
	fmt.Printf("10 - 3 = %d\n", subtract(10, 3))
	fmt.Printf("10 / 4 = %.1f\n", divide(10, 4))
}
