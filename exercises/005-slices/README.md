# Exercise 005 — Slices

This exercise introduces slices, Go’s primary data structure for working with lists of values.

Slices are dynamic and can grow or shrink as needed.

## Concepts Covered

- Creating slices
- Accessing slice elements
- Looping through slices with `range`
- Adding elements with `append`

## Goals

In this exercise, the program should:

1. Create a slice of names
2. Print each name
3. Add another name to the slice
4. Print the updated slice

## Example Output

    Original names:
    Brad
    Alice
    Sam

    After append:
    Brad
    Alice
    Sam
    Jordan

## Run the Exercise

From the repository root:

```bash
go run exercises/005-slices/main.go