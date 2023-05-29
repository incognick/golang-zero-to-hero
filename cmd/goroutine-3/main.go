package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := map[string]int{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			numbers[fmt.Sprintf("%v", n)] = n
		}(i)
	}

	wg.Wait()

	fmt.Println(numbers)
}
