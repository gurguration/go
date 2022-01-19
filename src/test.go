package main

import (
	"fmt"
)

func main() {
	s := []int{1821, 3332, 83, 34, 6}
	for k, v := range s {
		fmt.Printf("This is key %v, This is V: %v\n", k, v)
	}
}
