package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", "And Guram")
	}
}

func sayMessage(msg, name string) {
	fmt.Println(msg, name)
}
