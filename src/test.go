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
	specialty  string `required max: "10"`
	companions string
}

type Specialist struct {
	Doctor
	isSpecialist bool
}

func main() {
	i := 5
	switch { //doesn't have to be boolean like in if's case
	case i >= 0 && i <= 2:
		fmt.Println("Is between 0 and 3")
		fallthrough // Fallthrough is logicless it will execute anyway
	case i <= 5 && i >= 3:
		fmt.Println("Is between 2 and 6")
		fallthrough
	case i == 10:
		fmt.Println("I is 10")
		fallthrough
	default:
		fmt.Println("another number")
	}
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
