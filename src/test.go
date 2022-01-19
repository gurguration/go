package main

import (
	"fmt"
)

func main() {
	var ms *myStruct = &myStruct{foo: 55}
	fmt.Println(ms)

}

type myStruct struct {
	foo int
}
