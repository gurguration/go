package main

import "fmt"

type Greeter struct {
	greeting string
	name     string
}

func main() {
	b := Greeter{
		greeting: "Hello",
		name:     "GO",
	}
	res := b.greet()
	fmt.Println(res)
}

func (g *Greeter) greet() (result string) {
	g.name = "New Name"
	result = ("The greeting is: " + g.greeting + " And the name: " + g.name)
	return
}
