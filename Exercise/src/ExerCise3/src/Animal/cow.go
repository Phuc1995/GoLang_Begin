package Animal

import "fmt"

type Cow struct {
	Name string
	Birth int
	Milk int
}

func NewCow(name string, birth int, milk int) *Cow{
	cow := &Cow{
		Name: name,
		Birth: birth,
		Milk: milk,
		}
	return cow
}

func (c Cow) Shout()  {
	fmt.Println("O O O ")
}

func (c Cow) GiveBirth(birth int)  {
	fmt.Printf("Cow just born : ", birth)
	c.Birth += birth
}

func (c Cow) GiveMilk(milk int ) {
	c.Milk += milk
}

func (c Cow) Statistic() {
	fmt.Printf("Total birth: %d and milk %d of Cow  ",c.Birth, c.Milk)
}