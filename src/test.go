package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(1 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
