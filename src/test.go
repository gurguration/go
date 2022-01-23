package main

import (
	"context"
	"fmt"
	"log"
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
	timeout := 5 * time.Second
	ctx2, _ := context.WithTimeout(cn, timeout)
	// fmt.Println(ctx2.Err())
	time.Sleep(3 * time.Second)
	fmt.Println(ctx2)
	fn(cn)
}

func fn(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		log.Print("After two Sec")
	case <-time.After(1 * time.Second):
		log.Print("After one Sec")
	case <-ctx.Done():
		log.Print("Boom context is done")
	}
}
