package main

import "fmt"

type House struct {
	area int
	door *Door
}

type Door struct {
	color string
}

type SmallApartment struct {
	smallApartment House
	}

type Person struct {
	name string
	apartment *SmallApartment
}

func (h *House)ShowData() string{
	return fmt.Sprint("I am a house, my area is: ",h.area)
}

func (h *House) GetArea() int {
	return h.area
}

func (h *House) SetArea(area int){
	h.area = area
}

func (d *Door) ShowData() string {
	return fmt.Sprint("I am a door, my color is: ",d.color)
}

func (d *Door) GetColor() string {
	return d.color
}

func (d *Door) SetColor(color string) {
	d.color = color
}

func (d *SmallApartment) SetArea(area int ) {
	d.smallApartment.area = area
}

func (p *Person) ShowdataHouse()  {

	fmt.Println("Name person: ",p.name)
	fmt.Println("Infor house: ",p.apartment.smallApartment.ShowData())
	fmt.Println("Info color: ",p.apartment.smallApartment.door.ShowData())
	p.apartment.SetArea(50)
	fmt.Println("Area of Apartment: ",p.apartment.smallApartment.area)
}

func main()  {
	door := Door{"red"}
	house := House{
		area:500,
		door: &door,
	}
	apartment := SmallApartment{house}
	p := &Person{name:"Phuc", apartment:&apartment}
	p.ShowdataHouse()


}