package main

import "fmt"

func main() {
	fmt.Println("slices")

	names := []string{
		"Nick",
		"Bob",
		"Frank",
		"Joe",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

	print(names)

	change(names)

	print(names)

}

func change(n []string) {
	n[2] = "Chris"
}

func print(names []string) {
	fmt.Printf("values=%v, %p, %d, %d\n", names, names, len(names), cap(names))
}
