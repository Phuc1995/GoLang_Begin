package session2

import (
	"fmt"
)

type Chicken struct {
	Name  string
	Birth int
}

func NewChicken(name string, birth int) *Chicken {
	chikend := &Chicken{name, birth}
	return chikend
}

func (m Chicken) Shout() {
	fmt.Printf("I'm Chickent : %s\n", m.Name)
}

func (m *Chicken) GiveBirth(birth int) {
	fmt.Printf("Chiken just give birth %d egg\n", birth)
	m.Birth += birth
}

func (m Chicken) GiveMilk(milk int) {}

func (m *Chicken) Statistic() {
	fmt.Printf("I'm CHICKEN %s \t\tI was gave birth Egg %d \n", m.Name, m.Birth)
}
