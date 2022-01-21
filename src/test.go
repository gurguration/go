package main

import (
	"fmt"
	"sync"
)

// WAIT GROUP

var wg = sync.WaitGroup{}

func main() {
	msg := "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}

//func sayHello() {
//	fmt.Println("Hello")
//}
