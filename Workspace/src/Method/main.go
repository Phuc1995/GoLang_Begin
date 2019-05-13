package main

import (
	"fmt"
	"image/color"
)

func main() {
	r := Rect{5, 10}
	fmt.Println("area: ", r.area())
	fmt.Println("area: ", r.perim())

	red := color.RGBA{255,0,0,255}
	c := ColorRect{Rect{length:10, width:5}, red}
	c.width = c.width + 5
	fmt.Println("area: ", c.area())
	fmt.Println("area: ", c.perim())

}

type Rect struct {
	length, width int
}

func (r Rect) area() int  {
	return r.length * r.width
}

func (r Rect) perim() int {
	return (r.length + r.width)*2
}

type ColorRect struct {
	Rect
	color color.RGBA
}