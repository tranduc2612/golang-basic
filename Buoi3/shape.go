package main

import "math"

type shape interface {
	dienTich() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) dienTich() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) dienTich() float64 {
	return s.length * s.length
}
