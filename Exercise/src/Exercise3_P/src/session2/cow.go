package session2

import "fmt"

type Cow struct {
	Name  string
	Birth int
	Milk  int
}

func NewCow(name string, birth int, milk int) *Cow {
	cow := &Cow{name, birth, milk}
	return cow
}

func (m Cow) Shout() {
	fmt.Printf("I'm Cow : %s\n", m.Name)
}

func (m *Cow) GiveBirth(birth int) {
	m.Birth += birth
	fmt.Printf("COW just give birth %d cow\n", birth)
}

func (m *Cow) GiveMilk(milk int) {
	m.Milk += milk
}

func (m Cow) Statistic() {
	fmt.Printf("I'm COW: %s \t\tI was gave birth: %d2 \t\tMilk : %d \n", m.Name, m.Birth, m.Milk)
}
