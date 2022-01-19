package main

import (
	"fmt"
)

func main() {
	a := 42
	b := &a
	a = a + 1

	fmt.Printf("Value A: %v Value B: %v\n", a, *b)
	fmt.Printf("address A: %v address B: %v\n", &a, b)
}
