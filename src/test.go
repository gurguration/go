package main

import (
	"fmt"
	"reflect"
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
	originalDoc := reflect.TypeOf(Doctor{specialty: "veryvery long name for the doc"})
	field, _ := originalDoc.FieldByName("specialty")
	fmt.Println(originalDoc)
	fmt.Println(field)
	fmt.Println(field.Tag)
	myDoc := Specialist{
		isSpecialist: true,
		Doctor:       Doctor{specialty: "Psyciatrist"},
	}
	fmt.Println(myDoc)
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
