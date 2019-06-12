package main

import "fmt"

type Rectangle struct {
	side1 float32
	side2 float32
	location Location
}

func (r *Rectangle) ToString() string{
	return fmt.Sprintf("Area Rectangle: %f - Perimetrer circle: %f", r.Area(), r.Perimeter())
}

func (r *Rectangle) Area() float32 {
	area := (r.side1 + r.side2) *2
	return area
}

func (r *Rectangle) Perimeter() float32 {
	perimetrer := r.side2 * r.side1
	return perimetrer
}