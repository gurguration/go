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
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("This is i: %v, and this is J: %v\n ", i, j)
	}
}
