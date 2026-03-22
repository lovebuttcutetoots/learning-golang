# Exercise 011 — Error Handling

This exercise introduces error handling in Go.

Go functions often return a value and an error. The caller is responsible for checking and handling that error.

## Concepts Covered

- Returning errors from functions
- The `(value, error)` pattern
- Checking errors with `if err != nil`

## Goals

In this exercise the program should:

1. Create a function that divides two numbers
2. Return an error if division by zero is attempted
3. Handle the error in `main()`

## Example Output

    Result: 5

    Error: cannot divide by zero

## Run the Exercise

From the repository root:

```bash
go run exercises/011-error-handling/main.go
```
Or from the exercise folder:

```bash
cd exercises/011-error-handling
go run .
```

## Key Learning

Errors are returned values in Go and must be handled explicitly.


---

## main.go

Create `exercises/011-error-handling/main.go`:

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
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
}

```