package main

import "fmt"

type House struct {
	Beds  int
	Baths int
	Size  int
}

func main() {
	fmt.Println("maps")

	m := map[string]*House{
		"Nick": {
			Beds:  4,
			Baths: 2,
			Size:  2500,
		},
	}

	m["Nick"].Baths = 3

	fmt.Printf("%+v\n", m["Nick"])

	delete(m, "Nick")

	fmt.Printf("%+v\n", m)
}
