package main

import (
	"fmt"
)

func main() {
	s := sum(1, 22, 33, 2, 3, 41, 94)
	fmt.Println("The sum is:", *s)
}

func sum(numbers ...int) *int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return &sum
}
