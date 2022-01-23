package main

import (
	"fmt"
	"log"

	"../car"
)

func main() {
	c := car.NewCar()
	fmt.Println(c.DisplayBrandName(2055))
	displayablaBrand, err := c.DisplayBrandName(2022)
	if err != nil {
		log.Fatal("Err")
	}
	log.Println(displayablaBrand)
}
