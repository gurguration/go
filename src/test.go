package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(1 * time.Millisecond)
}

//func sayHello() {
//	fmt.Println("Hello")
//}
