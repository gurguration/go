package main

import (
	"fmt"
)

func main() {
	greeting := "Hello Go"
	name := "And Guram!"
	sayMessage(&greeting, &name)
	fmt.Println(name)
}

func sayMessage(msg, name *string) {
	*name = "and Lisa"
	fmt.Println(*msg, *name)
}
