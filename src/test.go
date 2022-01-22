package main

import (
	"fmt"
	"math"
	"reflect"
)

var i interface{}

type shape interface {
	area() float64
	perimeter() float64
}

type Polygon interface {
	getSides() int
}
type rectangle struct {
	width, heith float64
	sides        int
}

type circle struct {
	radius float64
}

type polygon struct {
	sides int
}

func (p polygon) getSides() int {
	return p.sides * 4
}
func (c circle) perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) perimeter() float64 {
	return r.heith * r.width / 2
}

func (r rectangle) getSides() int {
	result := r.sides * 4
	return result
}
func (r rectangle) area() float64 {
	result := r.heith * r.width / 2
	// fmt.Println(result)
	return result
}

func Calculate(s interface{}) {
	myType := reflect.TypeOf(s)
	fmt.Println(myType.Name())
	fmt.Println(myType.Kind())
	fmt.Println(reflect.ValueOf(s))
	// fileinfo := myType.Field(0)
	// fmt.Println(fileinfo.Name)

	// result := s.area()
	// if rect, ok := s.(rectangle); ok {
	// 	fmt.Printf("Rect implements sides function also %T\n", rect)
	// 	fmt.Printf("Calculating sides: %v\n", rect.getSides())
	// }
	// rad := s.perimeter()
	// fmt.Printf("Radius is: %v\n", rad)
	// return result
	switch x := s.(type) {
	case circle:
		fmt.Printf("type of interface %T\n", x)
		fmt.Println(x.perimeter())
	case rectangle:
		fmt.Printf("type of interface %T\n", x)
		fmt.Println("Calling getSides()")
		x.getSides()
	default:
		fmt.Println("Don't know the type")

	}
}

func main() {
	rshape := rectangle{width: 1.2, heith: 33.22, sides: 9}
	circle := circle{radius: 23.3}
	Calculate(circle)
	Calculate(rshape)
	// fmt.Printf("Type and value of an instance: %T, %v\n", rshape, rshape)
	// fmt.Printf("Type and value of an instance: %T, %v\n", circle, circle)
	// rshape.sides := 5
	// sides := rshape.getSides()
	i = make(map[string]interface{})
	// i = true
	// fmt.Printf("I is of type %T\n", i)
}
