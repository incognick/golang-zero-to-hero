package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Name   string
	Street string `json:"street_name"`
	State  string
	City   string
	Zip    int
}

func main() {
	fmt.Println("structs")

	data, err := os.ReadFile("address.json")
	if err != nil {
		panic(err)
	}

	a := Address{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		panic(err)
	}

	fmt.Println(a)
}
