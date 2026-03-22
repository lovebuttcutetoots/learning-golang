# Exercise 010 — Interfaces

This exercise introduces interfaces in Go.

Interfaces define behavior. Any type that implements the required methods satisfies the interface automatically.

## Concepts Covered

- Defining interfaces
- Implementing interfaces implicitly
- Using interfaces as function parameters

## Goals

In this exercise the program should:

1. Define a `Speaker` interface
2. Create multiple types that implement the interface
3. Use a function that accepts the interface
4. Call the function with different types

## Example Output

    Hello, my name is Brad.
    Hello, my name is Alice.

## Run the Exercise

From the repository root:

```bash
go run exercises/010-interfaces/main.go