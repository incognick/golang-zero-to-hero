package main

import (
	"fmt"
)

func main() {
	fmt.Println("switch")

	day := "Monday"

	if day == "Monday" {
		fmt.Println("Start of the week!")
	} else if day == "Saturday" || day == "Sunday" {
		fmt.Println("Weekend")
	} else {
		fmt.Println("Weekday")
	}

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday":
		fmt.Println("Start of the week!")
		fallthrough
	default:
		fmt.Println("Weekday")
	}

}
