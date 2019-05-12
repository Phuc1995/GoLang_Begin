package main

import "fmt"

//Interface
//multiple interface
//Embed interface
//empty interface
type Animal interface {
	speak()
}

type Movement interface {
	move()
}

//Embed interface
type NextAnimal interface {
	Animal
	Movement
}

func goout(i interface{})  {
	fmt.Println(i)
}

type Dog struct {

}

func (d Dog) speak()  {
	fmt.Println("Gau Gau")
}

func (d Dog) move() {
	fmt.Println("Di bang 4 chan")
}

type data struct {
	index int
}

func main()  {
	var animal Animal
	animal = Dog{}
	animal.speak()

	dog := Dog{}

	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	fmt.Println("====================================================")
	dog1 := Dog{}
	var na NextAnimal = dog1
	na.speak()
	na.move()

	fmt.Println("====================================================")
	goout(10.12)

	data := data{123}
	goout(data)
}
