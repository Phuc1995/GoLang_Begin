package Animal

import "fmt"

type Chicken struct {
	Name string
	Birth int
}

func NewChicken(name string, birth int) *Chicken {
	chicken := &Chicken{
		Name:name,
		Birth:birth,
	}
	return chicken
}

func (c Chicken) Shout()  {
	fmt.Println("O O O ")
}

func (c Chicken) GiveBirth(birth int)  {
	fmt.Printf("Chicken just has : ", birth)
	c.Birth += birth
}

func (c Chicken) GiveMilk(milk int ) {

}

func (c Chicken) Statistic() {
	fmt.Println("Total egg of chicken is:  ",c.Birth)
}
