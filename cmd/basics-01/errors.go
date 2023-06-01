package main

import (
	"fmt"
	"os"
)

func errorsEntry() {
	result, err := addWithError(12, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func addWithError(a, b int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("a is < 0")
	}
	if a > 10 {
		return 0, fmt.Errorf("a is > 10")
	}

	return a + b, nil
}
