package main

import (
	"fmt"
	"log"
)

func main() {

	defer fmt.Println("hello go world")
	defer func() {
		if p := recover(); p != nil {
			log.Println("Error", p)
			fmt.Println("Recovered")
		}
	}()
	panic("program crashed")
	fmt.Println("This should be deferred")
	fmt.Println("Panic")
	fmt.Println("end")
}
