package main

import "fmt"

func main() {
	numbersChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbersChan <- i
		}
		close(numbersChan)
	}()

	for n := range numbersChan {
		fmt.Println(n)
	}

	fmt.Println("done")
}
