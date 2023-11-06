package main

import "fmt"

// Write a program that returns the numbers from 1 to 100 as a string except when:
// 1. For multiples of three, return “Fizz”
// 2. For the multiples of five, return “Buzz”
// 3. For numbers which are multiples of both three and five return “Fizz Buzz”

// To check your answer, run "go test ."
func check(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "Fizz Buzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", input)
}
