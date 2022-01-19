package main

import (
	"fmt"
)

func main() {
	i := 0
	fmt.Println("Broke outside loop")
outerLoop:
	for i < 10 {
		i++
		for j := 0; j < 10; j++ {
			// fmt.Printf("i IS: %v\n", i)
			fmt.Printf("i IS: %v and this is J: %v\n", i, j)
			if j == 3 {
				continue outerLoop
				// break outerLoop
			}
		}

	}
}
