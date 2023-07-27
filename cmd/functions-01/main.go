package main

import "fmt"

func main() {
	fmt.Println("functions-01")
	fmt.Println(sum(1, 1))
}

func sum(a, b int) (err error, s int) {
	defer func() {
		if err != nil {
			s = 5
		}
	}()

	s = a + b
	err = fmt.Errorf("some error")

	return
}
