package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, heith float64
}

func (r rectangle) perimeter() float64 {
	return r.heith * r.width / 2
}

func (r rectangle) area() float64 {
	result := r.heith * r.width / 2
	fmt.Println(result)
	return result
}

func Calculate(s shape) float64 {
	result := s.area()
	return result
}

func main() {
	rshape := rectangle{width: 1.2, heith: 33.22}
	Calculate(rshape)
}
