package main

import "fmt"

type transformFunc func(int, int) (int, int)

func main() {
	fmt.Println("functions-as-arguments")

	fmt.Println(swap(1, 2))
	fmt.Println(increment(1, 2))
	fmt.Println(transform(swap, 1, 2))
	fmt.Println(transform(increment, 1, 2))
}

func transform(f transformFunc, a, b int) (int, int) {
	return f(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func increment(a, b int) (int, int) {
	return a + 1, b + 1
}

func multiply(a, b int) int {
	return a * b
}
