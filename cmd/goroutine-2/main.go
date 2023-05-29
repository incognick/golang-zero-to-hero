package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printNumber(n)
		}(i)
	}

	wg.Wait()
}

func printNumber(i int) {
	fmt.Println(i)
}
