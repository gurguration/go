package main

import "fmt"

type Greeter struct {
	greeting string
	name     string
}

func main() {
	g := Greeter{
		greeting: "Hello",
		name:     "GO",
	}
	g.greet()
}

func (g Greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
