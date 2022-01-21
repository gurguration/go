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
	b.greet()
	fmt.Printf("The new name is: %v\n", b.name)
}

func (g *Greeter) greet() {
	g.name = "New Name"
	fmt.Println(g.greeting, g.name)
}
