package main

import "fmt"

func typesEntry() {
	name := "Nick"

	a := 15
	b := int(15)
	c := uint(15)

	d := true
	e := false

	f := float32(3.5687)

	// Avoids compiler errors
	fmt.Println(name, a, b, c, d, e, f)
}
