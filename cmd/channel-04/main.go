package main

import (
	"fmt"
	"time"
)

func main() {
	slow := make(chan struct{})
	fast := make(chan struct{})
	tick := time.Tick(100 * time.Millisecond)

	go slowFunc(slow)
	go fastFunc(fast)

	for {
		select {
		case <-slow:
			fmt.Println("goose!")
		case <-fast:
			fmt.Println("duck")
		case <-tick:
			fmt.Println(".")
		}
	}

	fmt.Println("done")
}

func slowFunc(s chan struct{}) {
	for {
		time.Sleep(1 * time.Second)
		s <- struct{}{}
	}
}

func fastFunc(s chan struct{}) {
	for {
		time.Sleep(250 * time.Millisecond)
		s <- struct{}{}
	}
}
