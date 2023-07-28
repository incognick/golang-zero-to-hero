package main

import "fmt"

func main() {
	fmt.Println("functions-closure")

	s := sum()
	s2 := sum()

	fmt.Println(s(1))
	fmt.Println(s(1))
	fmt.Println(s2(5))
	fmt.Println(s(5))
}

func sum() func(int) int {
	runningTotal := 0
	return func(a int) int {
		runningTotal += a
		return runningTotal
	}
}
