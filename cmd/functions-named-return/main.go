package main

import "fmt"

func main() {
	fmt.Println("functions-named-return")

	fmt.Println(sum(1, 1))
}

func sum(a, b int) (result int) {
	result = a + b
	return
}
