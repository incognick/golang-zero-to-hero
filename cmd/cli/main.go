package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Adding")

	left, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	right, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	sum := left + right

	fmt.Println(sum)
}
