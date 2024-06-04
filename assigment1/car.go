package main

type Car struct {
	registation_number string
}

func NewCar(registation_number string) *Car {
	c := Car{registation_number: registation_number}
	return &c
}
