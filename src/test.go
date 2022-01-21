package main

import "fmt"

func main() {
	fmt.Println("implementing integer incremeneter interface")
	inc := IntCounter(0)
	var intIncrementerInterface Incremeneter
	intIncrementerInterface = &inc
	intIncrementerInterface.Increment()
	intIncrementerInterface.Increment()
	result := intIncrementerInterface.Increment()
	fmt.Println(result)
}

type Incremeneter interface {
	Increment() int
}
type IntCounter int

func (num *IntCounter) Increment() int {
	*num++
	return int(*num)
}
