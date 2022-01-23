package car

import (
	"errors"
	"fmt"
)

type car struct {
	Brand  Brand
	Year   int
	secret string
}

type Brand struct {
	Name       string
	Reputation int
}

func NewCar() car {
	return car{
		Brand: Brand{
			Name:       "Renault",
			Reputation: 9,
		},
		Year:   1981,
		secret: "top secret",
	}
}

func (car car) DisplayBrandName(yearNow int) (string, error) {
	displayabla := fmt.Sprintf("%v %d", car.Brand.Name, yearNow-car.Year)
	if yearNow > 2040 {
		return "", errors.New("Error message")
	}
	return displayabla, nil
}
