package main

import (
	"fmt"
	"strconv"
)

const (
	_ = iota
	one
	two
	three
)

func main() {
	fmt.Printf("Enumerable constants %v  %v  %v\n", one, two, three)
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
	fmt.Printf("This is byte represantation of string %v, of Type: %T", b, b)
	//strings are immutable but can be concatenated
	//myString[2] = byte(105) This results in compile time error

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
