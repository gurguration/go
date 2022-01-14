package main

import (
	"fmt"
	"strconv"
)

func main() {
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
	fa = 3.14e23
	fa = 2.1e15
	fmt.Printf("%v, %T\n", fa, fa) // ieee754 Floating point standard 32bit or 64 bit
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
