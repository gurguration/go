package car

type Car struct {
	Brand  Brand
	Year   int
	secret string
}

type Brand struct {
	Name       string
	Reputation int
}

func NewCar() Car {
	return Car{
		Brand: Brand{
			Name:       "Renault",
			Reputation: 9,
		},
		Year:   1981,
		secret: "top secret",
	}
}
