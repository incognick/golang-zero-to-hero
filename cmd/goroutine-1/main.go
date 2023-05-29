package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("nick")
	go hello("bob")
	go hello("jane")

	time.Sleep(1 * time.Second)
}

func hello(name string) {
	fmt.Println("hello ", name)
}
