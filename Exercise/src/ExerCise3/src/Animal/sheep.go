package Animal

import "fmt"

type Sheep struct {
	Name string
	Birth int
	Milk int
}

func NewSheep(name string, birth int, milk int) *Sheep{
	sheep := &Sheep{
		Name: name,
		Birth: birth,
		Milk: milk,
	}
	return sheep
}

func (c Sheep) Shout()  {
	fmt.Println("O O O ")
}

func (c Sheep) GiveBirth(birth int)  {
	fmt.Printf("Cow just born : ", birth)
	c.Birth += birth
}

func (c Sheep) GiveMilk(milk int ) {
	c.Milk += milk
}

func (c Sheep) Statistic() {
	fmt.Printf("Total birth: %d and milk %d of Cow  ",c.Birth, c.Milk)
}