package main

import "fmt"

func main() {
	fmt.Println("functions-defer")

	hello("Nick")
	defer hello("Frank")
	defer hello("Joe")
	hello("Bob")
}

func hello(name string) {
	fmt.Println("Hello", name)
}
