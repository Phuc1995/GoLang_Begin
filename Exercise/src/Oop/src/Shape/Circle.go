package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
	location Location
}

func (c *Circle) ToString() string{
	return fmt.Sprintf("Area circle: %f - Perimetrer circle: %f",c.Area(),c.Perimeter())
}

func (c *Circle) Area() float32 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (c *Circle) Perimeter() float32 {
	perimetrer := 2*math.Pi*c.radius
	return perimetrer
}
