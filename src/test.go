package main

import (
	"context"
	"fmt"
	"time"
	// "strconv"
	// "../car"
)

func main() {
	cn := context.Background()
	_ = context.WithValue(cn, "myKey", 123)
	// t := time.Now().Add(1 * time.Second)
	// ctx2, _ := context.WithDeadline(cn, t)
	// deadline, ok := ctx2.Deadline()
	// fmt.Println(deadline)
	// log.Print(ok)
	timeout := 1 * time.Second
	ctx2, _ := context.WithTimeout(cn, timeout)
	fmt.Println(ctx2.Err())
	time.Sleep(2 * time.Second)
	fmt.Println(ctx2.Err())
}
