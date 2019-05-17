package main

import "fmt"

type HornSounder interface {
	SoundHorn()
}

type Car struct {
	sound string
}

func (c Car)SoundHorn()  {
	fmt.Println(c.sound)
}

type Bike struct {
	sound string
}

func (b Bike)SoundHorn()  {
	fmt.Printf(b.sound)
}

func main()  {
	var vehicles = []HornSounder{
		Car{"car"},
		Bike{"Ken ken"},
	}

	var i int  = 1
	for _,v := range vehicles{
		v.SoundHorn()
		fmt.Println("v: ",v)
		i++
		fmt.Println(i)
	}
}