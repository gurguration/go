package main

import (
	"fmt"
)

const (
	_   = iota      // Iota initializer must be constant
	one = 1 << iota // m
	two
	three
	four
)

const (
	t        = 3
	y        = 'a'
	u        = "string"
	i uint64 = 53 + t // constants can be interpolated unlike variables
)

const (
	readonly = 1 << iota
	x1
	service
	x2
	admin
)

type Doctor struct {
	number     int
	actorName  string
	companions string
}

func main() {
	myDoc := Doctor{
		number:     3,
		actorName:  "Buba Kikabidze",
		companions: "Bidzina Bidzinashvili",
	}
	fmt.Println(myDoc)
	fmt.Println(myDoc.actorName)
	anotherDoc := Doctor{}
	anotherDoc.number = 10
	anotherDoc.actorName = "George Clooney"
	fmt.Println(myDoc)
	fmt.Println(anotherDoc)

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
