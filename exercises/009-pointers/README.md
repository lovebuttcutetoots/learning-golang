# Exercise 009 — Pointers

This exercise introduces pointers and how they are used to modify data in Go.

So far, methods have worked on copies of data. Pointers allow functions and methods to modify the original value.

## Concepts Covered

- What a pointer is
- Passing values vs passing pointers
- Pointer receivers in methods

## Goals

In this exercise the program should:

1. Define a `Person` struct
2. Create a method that modifies the person's age
3. Demonstrate the difference between value and pointer behavior

## Example Output

    Original age: 42
    After birthday: 43

## Run the Exercise

From the repository root:

```bash
go run exercises/009-pointers/main.go