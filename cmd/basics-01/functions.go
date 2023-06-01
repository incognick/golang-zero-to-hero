package main

import "fmt"

func funcEntry() {
	fmt.Println("------\nfuncEntry\n-----")
	result := add(1, 2)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}
