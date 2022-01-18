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
	//myMap := make(map[string]int)
	myMap := map[string]int{
		"TBILISI":  100,
		"Istanbul": 2,
		"Izmir":    1,
		"Prague":   1,
		"Berlin":   0,
		"Warsaw":   1,
	}
	delete(myMap, "Warsaw")
	myMapCopy := myMap
	fmt.Println(myMapCopy)

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
