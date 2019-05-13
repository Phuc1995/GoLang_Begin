package main

import "fmt"

type Shaper interface {
	area() int
	perim() int
}

type Rect struct {
	length, width int
}

type Square struct {
	side int
}

func (r Rect) area() int {
	return r.length*r.width
}
func (r Rect) perim() int {
	return (r.length + r.width)*2
}

func (sq Square) area() int  {
	return sq.side * sq.side
}

func (sq Square) perim() int {
	return sq.side * 2
}

func main()  {
	var s Shaper
	r := Rect{length:5, width:3}
	s = Shaper(r)
	fmt.Println("Area: ", s.area())
	fmt.Printf("Perimeter: ", s.perim())

	q := Square{side:5}
	s = q
	fmt.Printf("Arean ", s.area())
	fmt.Println("Perimeter: ", s.perim())
}