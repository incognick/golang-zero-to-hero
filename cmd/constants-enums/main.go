package main

import "fmt"

type State int

const (
	StatePrep State = iota + 1
	StateCooking
	StateDone
)

type Color int

const (
	ColorGreen Color = iota
	ColorBlue
	ColorRed
)

func main() {
	fmt.Println("constants and enums")

	handleState(StatePrep)
	// handleState(ColorGreen)
}

func handleState(state State) {
	fmt.Println(state)
}
