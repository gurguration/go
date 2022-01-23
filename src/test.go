package main

import (
	"fmt"
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
	fmt.Println(e)
}
