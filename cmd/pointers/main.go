package main

import "fmt"

func main() {
	fmt.Println("pointers")

	a := 5

	fmt.Printf("Location of a: %p, a=%v\n", &a, a)

	b := &a
	*b = 3

	fmt.Println("")

	fmt.Printf("Location of a: %p, a=%v\n", &a, a)
	fmt.Printf("Location of b: %p, b=%v, *b=%v\n", &b, b, *b)
}
