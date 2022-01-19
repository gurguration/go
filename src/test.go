package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	fmt.Println(ms) //uninitialized nill pointer
	// fmt.Println(ms.foo) //Error nil pointer dereference
	ms = new(myStruct)
	fmt.Println(ms.foo) //* operator has lower precedence than dot (.)
}

type myStruct struct {
	foo int
}
