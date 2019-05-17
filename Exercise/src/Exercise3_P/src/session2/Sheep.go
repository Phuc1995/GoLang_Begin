package session2

import "fmt"

type Sheep struct {
	Name  string
	Birth int
	Milk  int
}

func NewSheep(name string, birth int, milk int) *Sheep {
	sheep := &Sheep{name, birth, milk}
	return sheep
}

func (m Sheep) Shout() {
	fmt.Printf("I'm Sheep : %s\n", m.Name)
}

func (m *Sheep) GiveBirth(birth int) {
	m.Birth += birth
	fmt.Printf("SHEEP just give birth %d sheep\n", birth)
}

func (m *Sheep) GiveMilk(milk int) {
	m.Milk += milk
}

func (m *Sheep) Statistic() {
	fmt.Printf("I'm SHEEP: %s \t\tI was gave birth: %d \t\tMilk : %d \n", m.Name, m.Birth, m.Milk)
}
