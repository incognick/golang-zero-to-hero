package main

import "fmt"

var (
	list1 = []int{1, 2, 3, 4, 5}
	list2 = []int{6, 7, 8, 9, 10}
)

func main() {
	numbersChan := make(chan int)

	go sum(list1, numbersChan)
	go sum(list2, numbersChan)

	total1 := <-numbersChan
	total2 := <-numbersChan

	fmt.Println("total = ", total1+total2)
}

func sum(list []int, resp chan int) {
	total := 0
	for _, l := range list {
		total = total + l
	}
	resp <- total
}
