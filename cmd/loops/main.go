package main

import "fmt"

func main() {
	fmt.Println("loops")

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(i, sum)
	}

	fmt.Println("End:", sum)

	for sum < 20 {
		sum++
	}
	fmt.Println("sum < 20", sum)

	for {
		fmt.Println("hello")
	}
}
