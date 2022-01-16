package main

import (
	"fmt"
	"strconv"
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
	const role byte = readonly | service | admin
	fmt.Printf("Print or'ed (|) value of role: %b\n", role)
	fmt.Printf("Typed Constants: %b %T\n", readonly, readonly)
	fmt.Printf("Typed Constants: %b %T\n", service, service)
	fmt.Printf("Typed Constants: %b %T\n", admin, admin)
	fmt.Printf("Is admin: %v\n", admin&role == admin)
	fmt.Printf("Enumerable constants %v  %v  %v  %v\n", one, two, three, four)
	i := 1
	booli := 1 == 2
	boolb := 2 == 2
	var zeroVar bool
	a := 10             //  1010
	b := 3              //  0011
	fmt.Println(a & b)  //  0010
	fmt.Println(a | b)  //  1011
	fmt.Println(a ^ b)  //  1001
	fmt.Println(a &^ b) //  0100
	// BIT Shifting
	bita := 8 // 1000
	fmt.Println("BIT Shifting")
	fmt.Println(bita << 3)
	fmt.Println(bita >> 3)
	fmt.Println("BIT Shifting Done")
	fmt.Printf("%v: Uninitialized variables are set to zero(false), %T\n", zeroVar, zeroVar)
	fmt.Printf("%v, %T\n", booli, booli)
	fmt.Printf("%v, %T\n", boolb, boolb)
	i = 34
	var my_name string = "\nGuram"
	fmt.Println(i)
	myLoop()
	fmt.Printf("%v, %T\n", my_name, my_name)
	fmt.Printf("%v, %T\n", strconv.Itoa(loopvar), string(loopvar))
	fa := 3.14
	cn := 10.9
	fa = 3.14e23
	fa = 2.1e15
	fmt.Printf("%v, %T\n", fa, fa) // ieee754 Floating point standard 32bit or 64 bit
	fmt.Printf("Float arithmetic operations: %v", (fa / cn))
	myString := "This is my string"
	//Strings are aliases for bytes
	fmt.Printf("\nConvert string alias 105 back to string %v", string(myString[2]))
	fmt.Printf("\nMy sTring: %v, of Type: %T\n", myString[2], myString[2])
	fmt.Printf("\n" + myString + " My name is Guram\n")
	bs := []byte(myString)
	fmt.Printf("This is bytes string %v\n\n", bs)
	//s32 := 'a'
	var s32 rune
	s32 = 'a'
	fmt.Printf("Runes cannot be converted to bytes they are literal characters\n")
	fmt.Printf("Rune: %v, Type: %T\n", s32, s32)
	fmt.Printf("uint32 utf-32 type strings are single quoted: %v, %T\n", s32, s32)
	fmt.Printf("This is byte represantation of string %v, of Type: %T\n", b, b)
	//strings are immutable but can be concatenated
	//myString[2] = byte(105) This results in compile time error
	grades := [3]int{95, 33, 65}
	fmt.Printf("Grades array %v\n", grades)
	var myArray [3]int
	fmt.Printf("This is my array %v\n", myArray)
	mySecondArray := [...]int{1, 2, 34, 5}
	fmt.Printf("My second array %v\n", mySecondArray)
	myThirdArray := [3]string{"my", "name", "is"}
	fmt.Printf("This is my third array %v\n", myThirdArray)
	fmt.Printf("Length of my array: %v and capacity: %v\n", len(myThirdArray), cap(myThirdArray))
	fmt.Printf("Printing my string's first element: %v\n", myThirdArray[1])
	myMadeArray := make([]string, 1, 50)
	fmt.Println(cap(myMadeArray))
	myArrayCopy := &myThirdArray       //pass pointer to copy
	myArrayCopyOfaCopy := myThirdArray //Slices are at different addresses when copy
	myThirdArray[2] = "guram"
	fmt.Printf("Copy of my Third array: %v\n", myArrayCopy)
	fmt.Printf("Copy of A Copy: %v\n", myArrayCopyOfaCopy)
	// array concatenation
	// myConcatedArray := slice(myArrayCopyOfaCopy[:], myThirdArray[2])
	var myNewArray [3]int = [...]int{1, 2, 3}
	fmt.Printf("Declaring array explicitly and assigning value: %v\n", myNewArray)
	fmt.Printf("Capacity of new array %v\n", cap(myNewArray))
	fmt.Printf("Type of new array %T\n", myNewArray)
	//myMadeArray = append(myMadeArray, "guram") //append function can take multiple
	//myMadeArray = append(myMadeArray, "zurab") //arguments. and also can use js like
	//myMadeArray = append(myMadeArray, "dachi") //sytnax [...]{1,2,3}
	//myMadeArray = append(myMadeArray, "manana")
	myMadeArray = append(myMadeArray, "manana", "dachi", "zurab", "guram")
	fmt.Printf("Printing my made array %v, %T\n", myMadeArray, myMadeArray)
	fmt.Printf("Printing lenght of my made array %v\n", len(myMadeArray))
	//join two slices together by spread operator
	myMadeArray = append(myMadeArray, []string{"another slice", "joined"}...)
	myMadeArray = append(myMadeArray, myThirdArray[:]...)
	fmt.Printf("This is slice joined by spread operator %v\n", myMadeArray)
	fmt.Printf("Pop element from the slice: %v", myMadeArray[2:len(myMadeArray)-1])
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
