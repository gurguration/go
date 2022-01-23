package main

import (
	"errors"
	"fmt"
	"log"
)

type MyError struct {
	extraInfo string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("got an error: %s", e.extraInfo)
}
func main() {
	var e error
	e = &MyError{extraInfo: "SSSRrrrra, BOOM!!!"}
	// fmt.Println(e)
	var c *MyError
	check1 := errors.As(e, &c)
	log.Println(check1)
	log.Println(&c)
	eWrapped := fmt.Errorf("Wrapped error %w", e)
	check2 := errors.Is(eWrapped, e)
	log.Print(check2)

}
