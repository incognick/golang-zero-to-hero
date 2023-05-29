package main

import "fmt"

func main() {
	numbersChan := make(chan int, 9)

	for i := 0; i < 10; i++ {
		numbersChan <- i
	}

	fmt.Println("done")
}
