package main

import (
	"errors"
	"fmt"
	"math"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot take square root of a negative number")
	}

	return math.Sqrt(x), nil
}

func main() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = sqrt(64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = sqrt(-16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
