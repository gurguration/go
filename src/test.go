package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	i = 34
	var my_name string = "\nGuram"
	fmt.Println(i)
	myLoop()
	fmt.Printf("%v, %T\n", my_name, my_name)
	fmt.Printf("%v, %T\n", strconv.Itoa(loopvar), string(loopvar))
}

var loopvar int = 223

func myLoop() {
	var (
		actorName = 1
	)
	fmt.Printf("%v ", actorName)
	for i := loopvar; i < loopvar; i++ {
		fmt.Printf("%v, %t ", i, i)
	}
}

var (
	actorName string = "Elisabeth"
	season    int    = 11
)
