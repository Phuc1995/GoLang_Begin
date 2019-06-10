package main

import (
	"fmt"
	"strings"
)

type IPerson interface {
	ShowData() string
}

type Door struct {
	color string
}

func (d *Door) ShowData() string {
	return fmt.Sprintf("I am a door , my color is %s \n", d.color)
}

func (d *Door) GetColor() string {
	return d.color
}

func (d *Door) SetColor(s string) {
	d.color = s
}

type House struct {
	area float64
	door *Door
}

func (h *House) GetDoor(color string) *Door {
	h.door = new(Door)
	h.door.SetColor(color)
	return h.door
}

func (h *House) GetArea() float64 {
	return h.area
}

func (h *House) SetArea(f float64) {
	h.area = f
}

func (h *House) ShowData() string {
	return fmt.Sprintf("I am a house, my area is %f m2 \n", h.area)
}

type SmallAparment struct {
	smallHouse House
}

func GetSmallAparment(color string) *SmallAparment {
	s := new(SmallAparment)
	s.smallHouse.SetArea(float64(50))
	s.smallHouse.GetDoor(color)
	return s
}

type Person struct {
	name      string
	aparments map[string]*SmallAparment
}

func (p *Person) ShowData() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("I am %s \n", p.name))
	for key, val := range p.aparments {
		b.WriteString(fmt.Sprintf("House's %s \t %s \t %s \n", key, val.smallHouse.ShowData(), val.smallHouse.door.ShowData()))
	}
	return b.String()
}

func PrintData(s IPerson) {
	fmt.Println(s.ShowData())
}

func main() {
	aparments := make(map[string]*SmallAparment)
	aparments["Villa1"] = GetSmallAparment("Yellow")
	aparments["Villa2"] = GetSmallAparment("Red")
	aparments["Villa3"] = GetSmallAparment("Blue")
	p := &Person{name: "Tan", aparments: aparments}
	PrintData(p)
}
