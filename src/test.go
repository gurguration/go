package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(1 * time.Millisecond)
}

//func sayHello() {
//	fmt.Println("Hello")
//}
