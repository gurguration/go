package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting")
	fmt.Println("Entering panicker")
	panicker()
	fmt.Println("After panicker")
	fmt.Println("Finished")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered from error")
			log.Println("Error was: ", p)
		}
	}()
	panic("something bad happended")
}
