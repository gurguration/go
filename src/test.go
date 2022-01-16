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

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	a1 := a
	a1[1] = 432
	fmt.Printf("print type of a:%T and value:%v\n", a, a)
	fmt.Printf("print type of a1:%T and value:%v\n", a1, a1)
	var b []int = []int{1, 2, 3, 4, 5}
	b1 := b
	b1[1] = 433
	fmt.Printf("print type of b:%T and value:%v\n", b, b)
	fmt.Printf("print type of b1:%T and value:%v\n", b1, b1)
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
