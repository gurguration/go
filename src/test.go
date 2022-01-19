package main

import (
	"fmt"
)

func main() {
	mymap := map[string]string{"myname": "guram"}
	myMapCopy := mymap
	myMapCopy["myname"] = "lisa"
	fmt.Println(mymap, myMapCopy)
}

type myStruct struct {
	foo int
}
