package main

import "fmt"

type Fruit interface {
	String() string
}

type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("A %s apple", a.Variety)
}
type Orange struct {
	Size string
}

func (o Orange) String() string {
	return fmt.Sprintf("A %s orange", o.Size)
}

func PrinFrut(fruit Fruit)  {
	fmt.Println("I have this fruit: ", fruit.String())
}

func main()  {
	apple := Apple{"Golden"}
	orange := Orange{"12"}

	PrinFrut(apple)
	PrinFrut(orange)
}



