package main

import "fmt"

func main() {
	i := 1
	i = 34
	var my_name string = "\nGuram"
	fmt.Println(i)
	myLoop()
	fmt.Printf("%v, %T", my_name, my_name)
}

var loopvar int = 10

func myLoop() {
	var (
		actorName = 1
	)
	fmt.Printf("%v ", actorName)
	for i := loopvar; i < loopvar; i++ {
		println(i)
	}
}

var (
	actorName string = "Elisabeth"
	season    int    = 11
)
